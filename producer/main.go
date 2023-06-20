package main

import (
	"context"
	"fmt"
	"github.com/gempir/go-twitch-irc/v4"
	_ "github.com/lib/pq"
	"github.com/segmentio/kafka-go"
	"os"
)

func main() {
	var username = os.Getenv("CHANNEL")
	var conn = initializeKafka(username)

	defer conn.Close()

	client := twitch.NewAnonymousClient()
	client.OnConnect(func() {
		println("Connected to", username, ".")
	})

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(message.Message)

		msg := kafka.Message{
			Value: []byte(fmt.Sprint(message.Message)),
		}
		err := conn.WriteMessages(context.Background(), msg)
		if err != nil {
			fmt.Println(err)
		}
	})

	client.Join(username)
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}

func initializeKafka(topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr: kafka.TCP("kakfa:9092"),
	}
}
