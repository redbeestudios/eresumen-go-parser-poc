package out

import (
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/domain"
)

type ResumenProducer interface {
	SaveResumen(header *domain.ResumenAEnviar, opCount int, id string) error
	Flush()
}
