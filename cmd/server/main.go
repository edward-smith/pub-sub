package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"github.com/edward-smith/pub-sub/internal/pubsub"
	"github.com/edward-smith/pub-sub/internal/routing"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	connstr := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(connstr)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	fmt.Printf("connection successful...")
	defer conn.Close()

	chan, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer chan.Close()

	err = pubsub.PublishJSON(
		chan,
		routing.ExchangePerilDirect,
		routing.PauseKey, 
		routing.PlayingState{
			IsPaused: true,
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	fmt.Println("Pause message sent!")
}
