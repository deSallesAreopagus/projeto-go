package main

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	topic := "test-topic"
	kafkaBroker := os.Getenv("KAFKA_BROKER")
	if kafkaBroker == "" {
		kafkaBroker = "localhost:9092"
	}
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaBroker,
		"group.id":          "grupo-teste",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
	defer c.Close()

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}
	fmt.Printf("Subscribed to topic: %s\n", topic)

	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			continue
		}
		fmt.Printf("Received message: %s on topic %s\n", string(msg.Value), *msg.TopicPartition.Topic)
	}
}
