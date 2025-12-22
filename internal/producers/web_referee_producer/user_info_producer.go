package web_referee_producer

import (
	"encoding/json"
	"time"

	"context"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"github.com/tengenatari/web-referee/internal/models"
)

type WebRefereeProducer struct {
	writer *kafka.Writer
	topic  string
}

func NewWebRefereeProducer(broker string, topic string) *WebRefereeProducer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(broker),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		MaxAttempts:  3,
		BatchSize:    100,
		BatchBytes:   1048576,
		BatchTimeout: 10 * time.Millisecond,
		RequiredAcks: kafka.RequireAll,
		Async:        false,
		Compression:  kafka.Snappy,
	}
	return &WebRefereeProducer{
		writer: writer,
		topic:  topic,
	}
}

func (p *WebRefereeProducer) ProduceUser(ctx context.Context, user models.User) error {

	val, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "failed to marshal user")
	}
	msg := kafka.Message{
		Value: val,
		Time:  time.Now(),
	}

	return p.writer.WriteMessages(ctx, msg)

}

func (p *WebRefereeProducer) Close() error {
	return p.writer.Close()
}
