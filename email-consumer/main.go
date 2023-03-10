package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	cmdcommons "github.com/redbeestudios/go-parser-poc/commons/cmd"
	"github.com/redbeestudios/go-parser-poc/commons/pkg"
	"github.com/redbeestudios/go-parser-poc/email-consumer/cmd"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	env, err := pkg.NewEnv("dev")
	config := cmdcommons.InitConfig(env).EmailConsumer
	deps := cmd.InitDependencies(config)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(config.Kafka.ConfigMap())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)

	err = c.SubscribeTopics(config.Kafka.Topics(), nil)

	run := true

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
				deps.MasterCardConsumer.ProcessMastercard(*e)
				_, err := c.StoreMessage(e)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%% Error storing offset after message %s:\n",
						e.TopicPartition)
				}
			case kafka.Error:
				// Errors should generally be considered
				// informational, the client will try to
				// automatically recover.
				// But in this example we choose to terminate
				// the application if all brokers are down.
				fmt.Fprintf(os.Stderr, "%% Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()
}
