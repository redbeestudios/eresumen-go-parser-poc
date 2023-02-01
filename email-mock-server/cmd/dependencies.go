package cmd

import (
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	redis "github.com/redbeestudios/go-parser-poc/email-receiver/internal/adapter/redis/mastercard"
)

type Dependencies struct {
	MailRepository *redis.CachedMailAdapter
}

func InitDependencies(config *pkg.EmailMockServerConfig) *Dependencies {

	repository := redis.NewCachedMailAdapter(config)

	return &Dependencies{
		MailRepository: repository,
	}

}
