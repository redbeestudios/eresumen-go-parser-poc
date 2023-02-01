package in

import (
	"context"
	"github.com/redbeestudios/go-parser-poc/email-consumer/internal/application/domain"
)

//go:generate mockgen -source=./get_by_name.go -package=mocks -destination=../../../../mocks/get_pokemon_by_name.go

type SaveResumen interface {
	ProcessMastercard(ctx context.Context, mail *domain.Mail) error
}
