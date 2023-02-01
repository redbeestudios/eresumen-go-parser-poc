package usecase

import (
	"github.com/alitto/pond"
	"github.com/puzpuzpuz/xsync"
	. "github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/usecase/line_parser"
	"log"
	"time"
)

type SummaryBuilder struct {
}

func (c *SummaryBuilder) buildSummaries(accountNumbers *xsync.MapOf[string, []string]) *xsync.MapOf[string, *Resumen] {
	pool := pond.New(100, 2000)

	resumenesParciales := xsync.NewMapOf[*Resumen]()

	operationStart := time.Now()

	accountNumbers.Range(func(key string, value []string) bool {
		resumenACrear := NewResumen()
		resumenACrear.SetClave(ClaveFromString(value[0]))
		resumenesParciales.Store(key, resumenACrear)
		pool.Submit(func() {
			resumenACrear, _ = resumenesParciales.Load(key)
			for _, line := range value {
				resumenACrear = line_parser.AddFieldFromLine(line[0:5], line[151:], resumenACrear)
			}
			resumenesParciales.Store(key, resumenACrear)
		})
		return true
	})
	pool.StopAndWait()

	log.Printf("Operation Time: %s", time.Since(operationStart))
	return resumenesParciales
}

func NewSummaryBuilder() *SummaryBuilder {
	return &SummaryBuilder{}
}
