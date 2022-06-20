package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":	os.Getenv("KafkaBootstrapServers"),
	}

	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		log.Fatalf("Error producing" + err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *ckafka.Producer) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic: &topic,
			Partition: ckafka.PartitionAny,
		},
		Value: []byte(msg),
	}

	return producer.Produce(message, nil)
}