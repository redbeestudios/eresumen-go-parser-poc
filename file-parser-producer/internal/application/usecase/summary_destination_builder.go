package usecase

import (
	"github.com/puzpuzpuz/xsync"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/domain"
	"os/exec"
	"sync"
)

type SummaryDestinationBuilder struct{}

func (c *SummaryDestinationBuilder) iterateThroughSummaries(destinations *xsync.MapOf[string, *domain.ResumenAEnviar],
	summaries *xsync.MapOf[string, *domain.Resumen]) *domain.Resultado {

	uuid, _ := exec.Command("uuidgen").Output()

	resultado := domain.Resultado{OpCount: 0, ProcessId: string(uuid)}
	completeSummaries := xsync.NewMapOf[*domain.ResumenAEnviar]()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		destinations.Range(func(key string, value *domain.ResumenAEnviar) bool {
			summary, exists := summaries.Load(key)
			destination := value
			if !exists {
				resultado.DestSinLiq = append(resultado.DestSinLiq, key)
				return true
			}
			destination.SetResumen(summary)
			completeSummaries.Store(key, destination)
			resultado.OpCount = resultado.OpCount + 1
			return true
		})
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		summaries.Range(func(key string, value *domain.Resumen) bool {
			_, exists := destinations.Load(key)
			if !exists {
				resultado.LiqSinDest = append(resultado.LiqSinDest, key)
			}
			return true
		})
	}()

	wg.Wait()

	resultado.ResumenesAEnviar = completeSummaries

	return &resultado
}

func NewSummaryDestinationBuilder() *SummaryDestinationBuilder {
	return &SummaryDestinationBuilder{}
}
