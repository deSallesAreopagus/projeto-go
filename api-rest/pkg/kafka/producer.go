package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer

func InitProducer(broker string) {
	var err error
	producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		panic(err)
	}
}

func SendMessage(topic string, message string) error {
	if producer == nil {
		return fmt.Errorf("producer not initialized")
	}

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}

	return producer.Produce(msg, nil)
}

func CloseProducer() {
	if producer != nil {
		producer.Close()
	}
}
