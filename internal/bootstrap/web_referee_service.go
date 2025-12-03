package bootstrap

import (
	"github.com/tengenatari/web-referee/config"
	"github.com/tengenatari/web-referee/internal/services/web_referee_service"
	"github.com/tengenatari/web-referee/internal/storage/pgstorage"
)

func InitWebRefereeService(storage *pgstorage.PGstorage, cfg *config.Config) *web_referee_service.WebRefereeService {
	return web_referee_service.NewWebRefereeService(storage)
}
