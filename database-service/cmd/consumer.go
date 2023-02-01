package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alitto/pond"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/puzpuzpuz/xsync"
	cmdcommons "github.com/redbeestudios/go-parser-poc/commons/cmd"
	dto "github.com/redbeestudios/go-parser-poc/commons/dto/mastercard"
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartConsumer(config *pkg.Config, dependencies *Dependencies) {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(config.DatabaseService.Kafka.ConfigMap())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)

	err = c.SubscribeTopics(config.DatabaseService.Kafka.Topics(), nil)

	run := true

	pool := pond.New(100, 1000)

	processes := xsync.NewMapOf[*cmdcommons.Process]()

	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				ctx := context.Background()
				resumen, err := getMessage(e)
				process, exists := processes.Load(resumen.Id)
				if !exists {
					process = &cmdcommons.Process{
						StartTime: time.Now(),
						Count:     &xsync.Counter{},
					}
					processes.Store(resumen.Id, process)
				}
				dependencies.Consumer.SaveResumen(ctx, resumen.Resumen)
				process.Count.Add(1)

				if process.Count.Value() == int64(resumen.OpCount) {
					dependencies.Consumer.FinishSaving()
					log.Printf("Todos los mensajes procesados, total: %d", process.Count.Value())
					log.Printf("Tiempo total de ejecución: %s", time.Since(process.StartTime))
				}

				_, err = c.StoreMessage(e)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%% Error storing offset after message %s:\n",
						e.TopicPartition)
				}

			default:
			}
		}
	}
	pool.StopAndWait()

	fmt.Printf("Closing consumer\n")
	c.Close()

}

func getMessage(message *kafka.Message) (*dto.Mensaje, error) {
	resumen := &dto.Mensaje{}

	err := json.Unmarshal(message.Value, resumen)

	if err != nil {
		return nil, err
	}
	return resumen, nil
}
