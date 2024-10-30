package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream()

	// Create a durable consumer (subscription)
	_, err := js.Subscribe("events.*", func(msg *nats.Msg) {
		fmt.Printf("Received message: %s\n", string(msg.Data))
		if err := msg.Ack(); err != nil {
			log.Fatalf("Error acknowledging message: %v", err)
		}
	}, nats.Durable("EVENTS-DURABLE"), nats.ManualAck())
	if err != nil {
		log.Fatalf("Error subscribing to messages: %v", err)
	}

	// Keep the connection alive
	select {}
}
