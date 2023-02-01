package in

import (
	"context"
)

//go:generate mockgen -source=./get_by_name.go -package=mocks -destination=../../../../mocks/get_pokemon_by_name.go

type SaveResumen interface {
	ProcessMastercard(ctx context.Context, path string) error
}
