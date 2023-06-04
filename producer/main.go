package main

import (
	"context"
	"fmt"
	"github.com/gempir/go-twitch-irc/v4"
	_ "github.com/lib/pq"
	"github.com/segmentio/kafka-go"
	"log"
)

var username = "" // Twitch channel here

func main() {

	var conn = initializeKafka(username)

	client := twitch.NewAnonymousClient()
	client.OnConnect(func() {
		println("Connected to", username, ".")
	})

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)

		_, err := conn.WriteMessages(
			kafka.Message{Value: []byte(message.Message)},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

	})

	client.Join(username)
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

func initializeKafka(topic string) *kafka.Conn {

	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	println("Connected to Kafka.")

	return conn
}
