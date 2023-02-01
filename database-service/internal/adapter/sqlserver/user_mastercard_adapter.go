package sqlserver

import (
	"context"
	"github.com/puzpuzpuz/xsync"
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	dto "github.com/redbeestudios/go-parser-poc/database-service/internal/adapter/sqlserver/dto"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/port/out"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var _ out.UsuarioRepository = (*MastercardSqlUserAdapter)(nil)

type MastercardSqlUserAdapter struct {
	database *gorm.DB
	users    *xsync.MapOf[string, *dto.Usuario]
}

func (m MastercardSqlUserAdapter) Flush() error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	var usuarios []*dto.Usuario

	m.users.Range(func(key string, value *dto.Usuario) bool {
		usuarios = append(usuarios, value)
		return true
	})

	result := m.database.CreateInBatches(usuarios, 100)
	if result.Error != nil {
		return result.Error
	}

	m.users = xsync.NewMapOf[*dto.Usuario]()

	return nil
}

func (m MastercardSqlUserAdapter) SaveUsuario(usuario *domain.Usuario) error {

	userDto := &dto.Usuario{
		Nombre:          usuario.Nombre(),
		NumeroCuenta:    usuario.NumeroCuenta(),
		Direccion:       usuario.Direccion(),
		Localidad:       usuario.Localidad(),
		CodigoPostal:    usuario.CodigoPostal(),
		Email:           usuario.Email(),
		NumeroDocumento: usuario.NumeroDocumento(),
	}

	m.users.Store(usuario.NumeroCuenta(), userDto)

	return nil
}

func (m MastercardSqlUserAdapter) GetUsuario(nroCuenta string) (*domain.Usuario, error) {
	usuarioDto, err := m.getUsuario(nroCuenta)
	if err != nil {
		log.Printf("Could not find usuario %s error: %s", nroCuenta, err)
		return nil, err
	}

	usuario := usuarioDtoToDomain(usuarioDto)

	return usuario, nil
}

func (m *MastercardSqlUserAdapter) getUsuario(nroCuenta string) (*dto.Usuario, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 6*time.Second)
	defer cancel()

	var user *dto.Usuario = nil

	if result := m.database.Where("numero_cuenta = ?", nroCuenta).First(&user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func usuarioDtoToDomain(indexDto *dto.Usuario) *domain.Usuario {
	usuario := domain.NewUsuario()

	usuario.SetEmail(indexDto.Email)
	usuario.SetNombre(indexDto.Nombre)
	usuario.SetLocalidad(indexDto.Localidad)
	usuario.SetNumeroCuenta(indexDto.NumeroCuenta)
	usuario.SetNumeroDocumento(indexDto.NumeroDocumento)
	usuario.SetCodigoPostal(indexDto.CodigoPostal)
	usuario.SetDireccion(indexDto.Direccion)
	return usuario
}

func NewUserRepository(config *pkg.DatabaseServiceConfig) *MastercardSqlUserAdapter {

	sqldbConfig := config.Sql

	db, err := gorm.Open(sqlserver.Open(sqldbConfig.Uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	db.Debug().AutoMigrate(&dto.Usuario{})

	if err != nil {
		panic(err)
	}
	return &MastercardSqlUserAdapter{
		database: db,
		users:    xsync.NewMapOf[*dto.Usuario](),
	}
}
