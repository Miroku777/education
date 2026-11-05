package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaAddress = "localhost:9092"
	topic        = "my-topic"
	group        = "MyGroup"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaAddress},
		Topic:   topic,
		GroupID: group,
	})
	defer r.Close()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	readMessagesFromKafka(ctx, r)
}

func readMessagesFromKafka(ctx context.Context, r *kafka.Reader) {
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("I got message: %s from topic: %s\n", string(msg.Value), msg.Topic)
	}
}

//docker exec -it kafka /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --create --topic my-topic --partitions 1 --replication-factor 1
//-создание топика Created topic my-topic.
