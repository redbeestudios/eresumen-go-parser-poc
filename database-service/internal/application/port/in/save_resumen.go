package in

import (
	"context"
	"github.com/redbeestudios/go-parser-poc/commons/dto/mastercard"
)

//go:generate mockgen -source=./get_by_name.go -package=mocks -destination=../../../../mocks/get_pokemon_by_name.go

type SaveResumen interface {
	SaveResumen(ctx context.Context, resumen *mastercard.ResumenAEnviar) error
}
