package web_referee_service

import (
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

type WebRefereeStorage interface {
	CreateUser(ctx context.Context, user *models.User) error
	CreateTournament(ctx context.Context, tournament *models.Tournament) error
	CreatePlayer(ctx context.Context, player *models.Player) error
}

type WebRefereeProducer interface {
	ProduceUser(ctx context.Context, user models.User) error
}

type WebRefereeService struct {
	webRefereeStorage  WebRefereeStorage
	webRefereeProducer WebRefereeProducer
}

func NewWebRefereeService(webRefereeStorage WebRefereeStorage, producer WebRefereeProducer) *WebRefereeService {
	return &WebRefereeService{
		webRefereeStorage:  webRefereeStorage,
		webRefereeProducer: producer,
	}
}
