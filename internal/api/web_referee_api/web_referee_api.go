package web_referee_service_api

import (
	"github.com/google/uuid"
	"github.com/tengenatari/web-referee/internal/models"
	"github.com/tengenatari/web-referee/internal/pb/web_referee_api"
	"golang.org/x/net/context"
)

type WebRefereeService interface {
	HealthCheck(ctx context.Context) error
	CreateUser(ctx context.Context, user *models.User) (string, error)
	CreateTournament(ctx context.Context, tournament *models.Tournament) (string, error)
	CreatePlayer(ctx context.Context, player *models.Player) (string, error)
	CreatePairing(ctx context.Context, tournamentId uuid.UUID) error
	GetPairing(ctx context.Context, tournamentId uuid.UUID) ([]*models.Game, error)
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
