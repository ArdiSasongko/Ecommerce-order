package external

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ArdiSasongko/Ecommerce-order/internal/config/env"
	"github.com/ArdiSasongko/Ecommerce-order/internal/config/logger"
	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
)

var log = logger.NewLogger()

type KafkaExternal struct {
}

func (k *KafkaExternal) ProduceKafkaMessage(ctx context.Context, data []byte) error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second

	brokers := strings.Split(env.GetEnvString("KAFKA_BROKER", ""), ",")
	topic := env.GetEnvString("KAFKA_TOPIC", "")

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return fmt.Errorf("failed to create comunacation with kafka :%w", err)
	}
	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(data),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		return fmt.Errorf("failed to produce message kafka :%w", err)
	}

	log.Infof("successfully produce on topic %s, partition %d, offset %d", topic, partition, offset)
	return nil
}
