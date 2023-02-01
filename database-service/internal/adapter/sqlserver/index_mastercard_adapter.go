package sqlserver

import (
	"context"
	"github.com/puzpuzpuz/xsync"
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	dto "github.com/redbeestudios/go-parser-poc/database-service/internal/adapter/sqlserver/dto"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/port/out"
	sf "github.com/sa-/slicefunk"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var _ out.IndexRepository = (*MastercardSqlIndexAdapter)(nil)

type MastercardSqlIndexAdapter struct {
	database *gorm.DB
	indexes  *xsync.MapOf[string, *dto.IndiceResumen]
}

func (m MastercardSqlIndexAdapter) Flush() error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	var indices []*dto.IndiceResumen

	m.indexes.Range(func(key string, value *dto.IndiceResumen) bool {
		indices = append(indices, value)
		return true
	})

	result := m.database.CreateInBatches(
		indices,
		100,
	)
	if result.Error != nil {
		return result.Error
	}

	m.indexes = xsync.NewMapOf[*dto.IndiceResumen]()

	return nil
}

func (m MastercardSqlIndexAdapter) SaveIndex(index *domain.Indice, usuario string) error {
	indexDto := &dto.IndiceResumen{
		Usuario:       usuario,
		FechaCierre:   index.FechaCierre(),
		Marca:         index.Marca(),
		Path:          index.Path(),
		Month:         index.Month(),
		Year:          index.Year(),
		ModeloResumen: index.ModeloResumen(),
	}

	m.indexes.Store(usuario+index.FechaCierre(), indexDto)

	return nil
}

func (m MastercardSqlIndexAdapter) GetIndex(nroCuenta string, periodo string) (*domain.Indice, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	var indice *dto.IndiceResumen = nil

	if result := m.database.Where("usuario = ? AND fecha_cierre = ?", nroCuenta, periodo).First(&indice); result.Error != nil {
		return nil, result.Error
	}

	return indiceDtoToDomain(indice), nil
}

func (m MastercardSqlIndexAdapter) GetIndices(nroCuenta string) ([]*domain.Indice, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	var indice = []*dto.IndiceResumen{}

	if result := m.database.Where("usuario = ? AND fecha_cierre = ?", nroCuenta).Find(&indice); result.Error != nil {
		return nil, result.Error
	}

	indices := sf.Map[*dto.IndiceResumen, *domain.Indice](indice, func(item *dto.IndiceResumen) *domain.Indice {
		return indiceDtoToDomain(item)
	})

	return indices, nil
}

func indiceDtoToDomain(indexDto *dto.IndiceResumen) *domain.Indice {
	index := domain.NewIndice()

	index.SetPath(indexDto.Path)
	index.SetYear(indexDto.Year)
	index.SetMonth(indexDto.Month)
	index.SetMarca(indexDto.Marca)
	index.SetFechaCierre(indexDto.FechaCierre)
	index.SetModeloResumen(indexDto.ModeloResumen)
	return index
}

func NewIndexSqlRepository(config *pkg.DatabaseServiceConfig) *MastercardSqlIndexAdapter {

	sqldbConfig := config.Sql

	db, err := gorm.Open(sqlserver.Open(sqldbConfig.Uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	db.Debug().AutoMigrate(&dto.IndiceResumen{})

	if err != nil {
		panic(err)
	}
	return &MastercardSqlIndexAdapter{
		database: db,
		indexes:  xsync.NewMapOf[*dto.IndiceResumen](),
	}
}
