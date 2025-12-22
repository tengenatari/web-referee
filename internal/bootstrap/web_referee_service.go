package bootstrap

import (
	"github.com/tengenatari/web-referee/internal/producers/web_referee_producer"
	"github.com/tengenatari/web-referee/internal/services/web_referee_service"
	"github.com/tengenatari/web-referee/internal/storage/pgstorage"
)

func InitWebRefereeService(storage *pgstorage.WebRefereeStorage, producer *web_referee_producer.WebRefereeProducer) *web_referee_service.WebRefereeService {
	return web_referee_service.NewWebRefereeService(storage, producer)
}
