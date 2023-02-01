package kafka

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/domain"
	"github.com/redbeestudios/go-parser-poc/file-parser-producer/internal/application/port/out"
	"log"
)

var _ out.ResumenProducer = (*MasterCardProducer)(nil)

type MasterCardProducer struct {
	producer *kafka.Producer
	topic    string
}

func (a *MasterCardProducer) Flush() {
	a.producer.Flush(15 * 1000)
}

func NewKafkaProducer(config *pkg.FileParserProducerConfig) out.ResumenProducer {

	p, err := kafka.NewProducer(config.Kafka.ConfigMap())
	if err != nil {
		panic(err)
	}
	return &MasterCardProducer{producer: p, topic: config.Kafka.Topic}
}

func (a *MasterCardProducer) SaveResumen(resumen *domain.ResumenAEnviar, opCount int, id string) error {
	body, err := json.MarshalIndent(Serialize(resumen, opCount, id), "", " ")
	if err != nil {
		log.Printf("Error creando archivo para resumen %d", resumen.Destinatario().NumeroCuenta())
		return err
	}
	err = a.producer.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &a.topic, Partition: kafka.PartitionAny}, Value: body}, nil)
	if err != nil {
		log.Printf("Error creando archivo para resumen %d", resumen.Destinatario().NumeroCuenta())
	}

	return nil
}
