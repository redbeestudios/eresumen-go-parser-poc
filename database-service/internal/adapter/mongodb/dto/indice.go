package mongodb

import (
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IndiceResumen struct {
	Id            primitive.ObjectID `bson:"_id"`
	Usuario       string             `bson:"usuario,omitempty"`
	FechaCierre   string             `bson:"fecha_cierre,omitempty"`
	Marca         string             `bson:"marca,omitempty"`
	Path          string             `bson:"path,omitempty"`
	Month         string             `bson:"month,omitempty"`
	Year          string             `bson:"year,omitempty"`
	ModeloResumen string             `bson:"modelo_resumen,omitempty"`
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
