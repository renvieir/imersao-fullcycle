package kafka

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

type KakfaConsumer struct {
	MsgChan	chan* ckafka.Message
}

func NewKakfaConsumer(msgChan chan* ckafka.Message) *KakfaConsumer {
	return &KakfaConsumer{
		MsgChan: msgChan,
	}
}

func (kc *KakfaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":	os.Getenv("KafkaBootstrapServers"),
		"group.id":						os.Getenv("KakfaConsumerGroupId"),
	}

	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error consuming kafka message:" + err.Error())
	}

	topics := []string{
		os.Getenv("KafkaReadTopic"),
	}
	c.SubscribeTopics(topics, nil)
	fmt.Println("Kafka consumer has been started")
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			kc.MsgChan <- msg
		}
	}
}
