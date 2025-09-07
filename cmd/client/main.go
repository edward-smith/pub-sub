package main

import (
	"fmt"
	"log"

	"github.com/edward-smith/pub-sub/internal/gamelogic"
	"github.com/edward-smith/pub-sub/internal/pubsub"
	"github.com/edward-smith/pub-sub/internal/routing"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril client...")

	username, err := gamelogic.ClientWelcome()
	if err != nil {
		fmt.Printf("Error welcoming client: %v\n", err)
		return
	}

	connstr := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(connstr)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	fmt.Printf("connection successful...")
	defer conn.Close()

	if err != nil {
		fmt.Printf("Error connecting to message broker: %v\n", err)
		return
	}
	defer conn.Close()

	_, _, err = pubsub.DeclareAndBind(conn, routing.ExchangePerilDirect, fmt.Sprintf("pause.%s", username), routing.PauseKey, 1)
	if err != nil {
		fmt.Printf("Error declaring and binding queue: %v\n", err)
		return
	}

	fmt.Printf("Welcome, %s!\n", username)
}
