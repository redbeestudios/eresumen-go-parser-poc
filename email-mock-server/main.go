package main

import (
	"errors"
	"fmt"
	"github.com/emersion/go-smtp"
	cmdcommons "github.com/redbeestudios/go-parser-poc/commons/cmd"
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	"github.com/redbeestudios/go-parser-poc/email-receiver/cmd"
	redis "github.com/redbeestudios/go-parser-poc/email-receiver/internal/adapter/redis/mastercard"
	"io"
	"log"
	"time"
)

// The Backend implements SMTP server methods.
type Backend struct {
	RedisRepo *redis.CachedMailAdapter
}

func (bkd *Backend) NewSession(_ *smtp.Conn) (smtp.Session, error) {
	return &Session{
		MailDto:  &redis.Mail{},
		MailRepo: bkd.RedisRepo,
	}, nil
}

// A Session is returned after EHLO.
type Session struct {
	MailDto  *redis.Mail
	MailRepo *redis.CachedMailAdapter
}

func (s *Session) AuthPlain(username, password string) error {
	if username != "username" || password != "password" {
		return errors.New("Invalid username or password")
	}
	return nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	s.MailDto.From = from
	return nil
}

func (s *Session) Rcpt(to string) error {
	s.MailDto.To = to
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := io.ReadAll(r); err != nil {
		return err
	} else {
		s.MailDto.Body = string(b)
		//s.MailRepo.Save(s.MailDto)
	}
	return nil
}

func (s *Session) Reset() {
}

func (s *Session) Logout() error {
	return nil
}

func main() {

	env, err := pkg.NewEnv("dev")
	if err != nil {
		panic(fmt.Sprintf("error creating environment: %s", err.Error()))
	}

	config := cmdcommons.InitConfig(env).EmailMockServer

	deps := cmd.InitDependencies(config)

	be := &Backend{
		RedisRepo: deps.MailRepository,
	}

	if config.Enabled {
		s := smtp.NewServer(be)

		s.Addr = fmt.Sprintf(":%s", config.Port)
		s.Domain = config.Address
		s.ReadTimeout = time.Duration(config.ReadTimeout) * time.Second
		s.WriteTimeout = time.Duration(config.WriteTimeout) * time.Second
		s.MaxMessageBytes = config.MaxMessageBytes * 10024
		s.MaxRecipients = config.MaxRecipients
		s.AllowInsecureAuth = config.AllowInsecureAuth

		log.Println("Starting server at", s.Addr)
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}

}
