package usecase

import (
	"context"
	"github.com/puzpuzpuz/xsync"
	. "github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/port/in"
	"log"
	"sync"
	"time"
)

var _ in.SaveResumen = (*SaveResumen)(nil)

type SaveResumen struct {
	summaryFileReader         *SummaryFileReader
	summaryBuilder            *SummaryBuilder
	summaryNotifier           *SummaryNotifier
	summaryDestinationReader  *SummaryDestinationReader
	summaryDestinationBuilder *SummaryDestinationBuilder
}

func (c *SaveResumen) ProcessMastercard(ctx context.Context, path string) error {

	start := time.Now()

	var wg sync.WaitGroup

	var resumenesParciales *xsync.MapOf[string, *Resumen]

	var resumenesAEnviar *xsync.MapOf[string, *ResumenAEnviar]

	var resultado *Resultado

	wg.Add(1)
	go func() {
		defer wg.Done()

		summaryLinesByType, err := c.summaryFileReader.readSummaryLines(path)
		if err != nil {
			return
		}

		log.Printf("File successfully read, all Records stored on memory, elapsed time: %s", time.Since(start))

		resumenesParciales = c.summaryBuilder.buildSummaries(summaryLinesByType)

		log.Printf("Summaries built, elapsed time: %s", time.Since(start))
	}()

	wg.Add(1)
	go func() {

		defer wg.Done()
		resumenesAEnviar, _ = c.summaryDestinationReader.getDestinations(path)

		log.Printf("Destinations loaded, elapsed time: %s", time.Since(start))

	}()

	wg.Wait()
	resultado = c.summaryDestinationBuilder.iterateThroughSummaries(resumenesAEnviar, resumenesParciales)

	c.summaryNotifier.notifyCreatedSummaries(resultado.ResumenesAEnviar, resultado.OpCount, resultado.ProcessId)

	log.Printf("All summaries notified, Elapsed time: %s", time.Since(start))

	log.Printf("Cuentas presentes en archivo de liquidacion pero no presentes en archivo de mails: %d:", len(resultado.LiqSinDest))

	log.Printf("Cuentas presentes en archivo de mails pero no presentes en archivo de liq: %d:", len(resultado.DestSinLiq))

	return nil
}

func NewSaveResumen(summaryFileReader *SummaryFileReader, summaryBuilder *SummaryBuilder, summaryNotifier *SummaryNotifier, summaryDestinationReader *SummaryDestinationReader, summaryDestinationBuilder *SummaryDestinationBuilder) *SaveResumen {
	return &SaveResumen{
		summaryFileReader:         summaryFileReader,
		summaryBuilder:            summaryBuilder,
		summaryNotifier:           summaryNotifier,
		summaryDestinationReader:  summaryDestinationReader,
		summaryDestinationBuilder: summaryDestinationBuilder,
	}
}
