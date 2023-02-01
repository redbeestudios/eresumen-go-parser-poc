package out

import (
	"github.com/redbeestudios/go-parser-poc/commons/models/resumen_mensual"
)

type ResumenRepository interface {
	SaveResumen(header *resumen_mensual.Resumen) error
}
