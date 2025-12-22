package bootstrap

import (
	"fmt"

	"github.com/tengenatari/web-referee/config"
	"github.com/tengenatari/web-referee/internal/producers/web_referee_producer"
)

func InitMessageProducer(cfg *config.Config) *web_referee_producer.WebRefereeProducer {
	broker := fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)
	return web_referee_producer.NewWebRefereeProducer(broker, cfg.Kafka.Topic)
}
