package out

import (
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
)

type IndexRepository interface {
	SaveIndex(index *domain.Indice, usuario string) error
	GetIndex(nroCuenta string, periodo string) (*domain.Indice, error)
	GetIndices(nroCuenta string) ([]*domain.Indice, error)
	Flush() error
}
