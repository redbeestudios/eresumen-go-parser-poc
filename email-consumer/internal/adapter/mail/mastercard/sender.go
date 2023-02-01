package mail

import (
	"fmt"
	"github.com/emersion/go-smtp"
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	"github.com/redbeestudios/go-parser-poc/email-consumer/internal/application/domain"

	"log"
	"strings"
)

type Sender struct {
	Config *pkg.EmailConsumerConfig
}

func NewMasterCardMailSender(config *pkg.EmailConsumerConfig) *Sender {
	return &Sender{
		Config: config,
	}

}

func (c *Sender) Send(message *domain.Mail) {
	client, err := smtp.Dial(c.Config.Smtp.Url)
	if err != nil {
		log.Println(err)
	}
	defer client.Close()

	// Set the sender and recipient, and send the email all in one step.
	to := []string{message.Email()}
	msg := strings.NewReader(fmt.Sprintf("%s", BuildMailBody(message)))
	err = client.SendMail("sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}

}
