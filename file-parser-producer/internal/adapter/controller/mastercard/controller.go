package mastercard

import (
	"context"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/port/in"
)

type Controller struct {
	saveHeaders in.SaveResumen
}

func NewMasterCardController(saveHeaders in.SaveResumen) *Controller {
	return &Controller{
		saveHeaders: saveHeaders,
	}
}

func (c *Controller) ProcessMastercard(ctx context.Context, path string) {
	go func() {
		c.saveHeaders.ProcessMastercard(ctx, path)
		return
	}()

}
