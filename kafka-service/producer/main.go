package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost: 9092"})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	topic := "test-topic"
	message := "Um teste de mensagens 2"

	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	e := <-p.Events()
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Error: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Success: \n Topic %s [%d] at offset %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
}
