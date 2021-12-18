package producer

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type service struct {
	topic string
	publisher *kafka.Producer
}

func NewProducer(p *kafka.Producer, topic string) *service {
	return &service{
		publisher: p,
		topic: topic,
	}
}

func (s *service) Push(data []byte) error {
	err := s.publisher.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &s.topic,
			Partition: kafka.PartitionAny,
		},
		Value: data,
	}, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to publish message in kafka")
		return err
	}
	return nil
}