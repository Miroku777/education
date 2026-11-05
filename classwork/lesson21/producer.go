package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaAddress = "localhost:9092"
	topic        = "my-topic"
)

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{kafkaAddress},
		Topic:   topic,
	})
	defer w.Close()
	messages := []string{
		"Hello",
		"Dear Students",
		"My Name is",
		"Dmitry",
	}
	sendMessagesToKafka(w, messages)
	log.Println("All messages were sent!")
}

func sendMessagesToKafka(w *kafka.Writer, messages []string) {
	for i, msg := range messages {
		newMessage := kafka.Message{
			Key:   []byte(fmt.Sprintf("Message #%d", i+1)),
			Value: []byte(msg),
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		err := w.WriteMessages(ctx, newMessage)
		if err != nil {
			log.Println(err)
			continue
		} else {
			log.Printf("Message #%d were sent", i)
		}
		time.Sleep(3 * time.Second)
	}
}
