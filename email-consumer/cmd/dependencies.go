package cmd

import (
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	kafka "github.com/redbeestudios/go-parser-poc/email-consumer/internal/adapter/kafka/mastercard"
	mail "github.com/redbeestudios/go-parser-poc/email-consumer/internal/adapter/mail/mastercard"
	"github.com/redbeestudios/go-parser-poc/email-consumer/internal/application/port/in"
	"github.com/redbeestudios/go-parser-poc/email-consumer/internal/application/usecase"
)

type Dependencies struct {
	MasterCardConsumer *kafka.MasterCardConsumer
}

func InitDependencies(config *pkg.EmailConsumerConfig) *Dependencies {

	saveHeaders := usecase.NewSaveResumen(mail.NewMasterCardMailSender(config))

	consumer := kafka.NewMasterCardConsumer(in.SaveResumen(saveHeaders))

	return &Dependencies{
		MasterCardConsumer: consumer,
	}

}
