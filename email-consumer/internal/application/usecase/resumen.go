package usecase

import (
	"context"
	mail "github.com/redbeestudios/go-parser-poc/email-consumer/internal/adapter/mail/mastercard"
	"github.com/redbeestudios/go-parser-poc/email-consumer/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/email-consumer/internal/application/port/in"
)

var _ in.SaveResumen = (*SaveResumen)(nil)

type SaveResumen struct {
	mailSender *mail.Sender
}

func (c *SaveResumen) ProcessMastercard(ctx context.Context, mail *domain.Mail) error {

	c.mailSender.Send(mail)

	return nil
}

func NewSaveResumen(mailSender *mail.Sender) *SaveResumen {
	return &SaveResumen{mailSender: mailSender}
}
