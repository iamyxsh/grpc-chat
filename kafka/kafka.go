package kafka

import (
	"context"
	kafka "github.com/segmentio/kafka-go"
	"log"
)

func ReturnConn(topic string) *kafka.Conn {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka:9092", topic, 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	return conn
}

func CreateTopics() {
	conn, err := kafka.Dial("tcp", "kafka:9092")
	if err != nil {
		log.Fatal("Kafka Connection error: ", err.Error())
	}
	defer conn.Close()

	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", "kafka:9092")
	if err != nil {
		log.Fatal("Kafka Dial error: ", err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             "USER_LOGIN",
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		log.Fatal("Kafka Topic Creation error: ", err.Error())
	}
}
