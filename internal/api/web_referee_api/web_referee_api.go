package web_referee_service_api

import (
	"github.com/tengenatari/web-referee/internal/models"
	"github.com/tengenatari/web-referee/internal/pb/web_referee_api"
	"golang.org/x/net/context"
)

type WebRefereeService interface {
	HealthCheck(ctx context.Context) error
	CreateUser(ctx context.Context, user *models.User) error
}

type WebRefereeServiceAPI struct {
	web_referee_api.UnimplementedWebRefereeServiceServer
	webRefereeService WebRefereeService
}

func NewWebRefereeServiceAPI(webRefereeService WebRefereeService) *WebRefereeServiceAPI {
	return &WebRefereeServiceAPI{
		webRefereeService: webRefereeService,
	}
}
