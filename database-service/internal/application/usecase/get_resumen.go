package usecase

import (
	"context"
	"errors"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/database-service/internal/application/port/out"
	"strings"
)

type GetResumen struct {
	indexRepository   out.IndexRepository
	resumenRepository out.ResumenRepository
}

func (c *GetResumen) GetResumen(ctx context.Context, codigoVinculacion string) (*domain.Resumen, error) {
	codSplit := strings.Split(codigoVinculacion, "-")

	cuenta := codSplit[0]

	periodo := codSplit[1]

	resumenIndex, err := c.indexRepository.GetIndex(cuenta, periodo)
	if err != nil {
		return nil, errors.New("No existe ese resumen para el usuario")
	}

	resumen, err := c.resumenRepository.GetResumen(resumenIndex.Path())
	if err != nil {
		return nil, errors.New("No existe ese resumen para el usuario")
	}

	return resumen, nil

}

func NewGetResumen(resumenRepository out.ResumenRepository, indexRepository out.IndexRepository) *GetResumen {
	return &GetResumen{
		resumenRepository: resumenRepository,
		indexRepository:   indexRepository,
	}
}
