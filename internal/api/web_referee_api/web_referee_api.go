package web_referee_service_api

import (
	"github.com/tengenatari/web-referee/internal/pb/web_referee_api"
	"golang.org/x/net/context"
)

type webRefereeService interface {
	HealthCheck(ctx context.Context) error
}

type WebRefereeServiceAPI struct {
	web_referee_api.UnimplementedWebRefereeServiceServer
	webRefereeService webRefereeService
}

func NewWebRefereeServiceAPI(webRefereeService webRefereeService) *WebRefereeServiceAPI {
	return &WebRefereeServiceAPI{
		webRefereeService: webRefereeService,
	}
}
