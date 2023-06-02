package kafka

import (
	"telebotV2/global"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

const (
	KafkaBroker = "209.159.148.138:9092"
	// KafkaBroker       = "localhost:9092"
)

type KafkaService struct {
	Writer *kafka.Writer
}

func NewKafkaService() *KafkaService {
	global.LOG.Info("kafka broker: ", zap.String("broker", KafkaBroker))
	writer := &kafka.Writer{
		Addr:     kafka.TCP(KafkaBroker),
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaService{
		Writer: writer,
	}
}
