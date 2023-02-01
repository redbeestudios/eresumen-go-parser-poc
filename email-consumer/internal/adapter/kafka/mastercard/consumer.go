package kafka

import (
	"context"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	dto "github.com/redbeestudios/go-parser-poc/commons/dto/mastercard"
	"github.com/redbeestudios/go-parser-poc/email-consumer/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/email-consumer/internal/application/port/in"
	"log"
)

type MasterCardConsumer struct {
	saveHeaders in.SaveResumen
}

func NewMasterCardConsumer(saveHeaders in.SaveResumen) *MasterCardConsumer {
	return &MasterCardConsumer{
		saveHeaders: saveHeaders,
	}
}

func (c *MasterCardConsumer) ProcessMastercard(message kafka.Message) {
	go func() {
		ctx := context.Background()

		var resumenMessage *dto.ResumenAEnviar

		err := json.Unmarshal(message.Value, &resumenMessage)

		if err != nil {
			log.Printf("Error unmarshalling msg")
			return
		}

		c.saveHeaders.ProcessMastercard(ctx, ToDomain(resumenMessage))

	}()

}

func ToDomain(r *dto.ResumenAEnviar) *domain.Mail {
	resumenAEnviar := domain.NewMail(
		r.Destinatario.Nombre,
		r.FechaCierre,
		r.CodigoVinculacion,
		r.Destinatario.Email,
	)

	return resumenAEnviar

}
