package usecase

import (
	"github.com/alitto/pond"
	"github.com/puzpuzpuz/xsync"
	. "github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/port/out"
	"log"
	"time"
)

type SummaryNotifier struct {
	resumenProducer out.ResumenProducer
}

func (c *SummaryNotifier) notifyCreatedSummaries(resumenesParciales *xsync.MapOf[string, *ResumenAEnviar], opCount int, id string) {
	operationStart := time.Now()

	pool := pond.New(100, 1000)

	resumenesParciales.Range(func(key string, value *ResumenAEnviar) bool {
		pool.Submit(func() {
			err := c.resumenProducer.SaveResumen(value, opCount, id)
			if err != nil {
				log.Printf("Error saving header in database: %s %s", value.Destinatario().NumeroCuenta(), err.Error())
			}
		})
		return true
	})
	pool.StopAndWait()
	c.resumenProducer.Flush()

	log.Printf("Operation Time: %s", time.Since(operationStart))
}

func NewSummaryNotifier(resumenProducer out.ResumenProducer) *SummaryNotifier {
	return &SummaryNotifier{
		resumenProducer: resumenProducer,
	}
}
