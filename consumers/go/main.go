package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
	"strings"
)

func main() {
	var username = os.Getenv("CHANNEL")

	r := getKafkaReader("kafka:9092", username)
	fmt.Println("Connected to Kafka...")
	defer r.Close()

	for {
		m, _ := r.ReadMessage(context.Background())
		fmt.Println(string(m.Value))
	}
}

func getKafkaReader(kafkaURL, topic string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		MinBytes: 10e1,
		MaxBytes: 10e6,
	})
}
