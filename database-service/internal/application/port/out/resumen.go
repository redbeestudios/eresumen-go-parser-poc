package out

import (
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
)

type ResumenRepository interface {
	SaveResumen(resumen *domain.Resumen, codVinculacion string) (string, error)
	GetResumen(codVinculacion string) (*domain.Resumen, error)
}
