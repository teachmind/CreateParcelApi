package publisher

import (
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func NewPublisher(host string) (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": host})
	if err != nil {
		return nil, err
	}
	return producer, nil
}
