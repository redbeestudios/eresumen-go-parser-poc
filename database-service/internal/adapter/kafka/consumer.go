package mastercard

import (
	"context"
	dto "github.com/redbeestudios/go-parser-poc/commons/dto/mastercard"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/adapter/serialization"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/port/in"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/usecase"
)

var _ in.SaveResumen = (*Consumer)(nil)

type Consumer struct {
	saveResumen *usecase.SaveResumen
}

func (c *Consumer) SaveResumen(ctx context.Context, resumen *dto.ResumenAEnviar) error {
	usuario, indice, resumenAGuardar := serialization.Deserialize(resumen)

	return c.saveResumen.SaveResumen(ctx, usuario, indice, resumenAGuardar)
}

func (c *Consumer) FinishSaving() error {
	return c.saveResumen.FinishSaving()
}

func NewMasterCardConsumer(saveResumen *usecase.SaveResumen) *Consumer {
	return &Consumer{
		saveResumen: saveResumen,
	}
}
