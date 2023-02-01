package cmd

import (
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/adapter/controller/mastercard"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/adapter/kafka"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/port/in"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/usecase"
)

type Dependencies struct {
	MasterCardController *mastercard.Controller
}

func InitDependencies(config *pkg.FileParserProducerConfig) *Dependencies {

	kafkaProducer := kafka.NewKafkaProducer(config)

	summaryFileReader := usecase.NewSummaryFileReader()
	summaryBuilder := usecase.NewSummaryBuilder()
	summaryNotifier := usecase.NewSummaryNotifier(kafkaProducer)
	summaryDestinationReader := usecase.NewSummaryDestinationReader()
	summaryDestinationBuilder := usecase.NewSummaryDestinationBuilder()

	saveHeaders := usecase.NewSaveResumen(summaryFileReader, summaryBuilder, summaryNotifier, summaryDestinationReader, summaryDestinationBuilder)

	controller := mastercard.NewMasterCardController(in.SaveResumen(saveHeaders))

	return &Dependencies{
		MasterCardController: controller,
	}

}
