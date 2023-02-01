package mastercard

import (
	"context"
	dto "github.com/redbeestudios/go-parser-poc/commons/dto/mastercard"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/adapter/serialization"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/port/in"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/usecase"
)

var _ in.GetResumen = (*Controller)(nil)
var _ in.SaveResumen = (*Controller)(nil)

type Controller struct {
	saveResumen *usecase.SaveResumen
	getResumen  *usecase.GetResumen
}

func NewMasterCardController(saveResumen *usecase.SaveResumen, getResumen *usecase.GetResumen) *Controller {
	return &Controller{
		saveResumen: saveResumen,
		getResumen:  getResumen,
	}
}

func (c *Controller) SaveResumen(ctx context.Context, resumen *dto.ResumenAEnviar) error {

	usuario, indice, resumenAGuardar := serialization.Deserialize(resumen)

	return c.saveResumen.SaveResumen(ctx, usuario, indice, resumenAGuardar)

}

func (c *Controller) GetResumen(ctx context.Context, codigoVinculacion string) (*dto.Resumen, error) {

	resumen, err := c.getResumen.GetResumen(ctx, codigoVinculacion)

	resumenAEnviar := serialization.SerializeResumen(resumen)

	return resumenAEnviar, err

}
