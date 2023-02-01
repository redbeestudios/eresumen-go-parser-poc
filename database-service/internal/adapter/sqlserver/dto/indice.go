package sqlserver

import (
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
	"gorm.io/gorm"
)

type IndiceResumen struct {
	gorm.Model
	Usuario       string `gorm:"column:usuario"`
	FechaCierre   string `gorm:"column:fecha_cierre"`
	Marca         string `gorm:"column:marca"`
	Path          string `gorm:"column:path"`
	Month         string `gorm:"column:month"`
	Year          string `gorm:"column:year"`
	ModeloResumen string `gorm:"column:modelo_resumen"`
}

func IndiceResumenFromDomain(resumen *domain.Indice, usuario string, path string) *IndiceResumen {
	return &IndiceResumen{
		Usuario:       usuario,
		FechaCierre:   resumen.FechaCierre(),
		Marca:         resumen.Marca(),
		Path:          path,
		Month:         resumen.Month(),
		Year:          resumen.Year(),
		ModeloResumen: resumen.ModeloResumen(),
	}
}

func (IndiceResumen) TableName() string {
	return "index"
}
