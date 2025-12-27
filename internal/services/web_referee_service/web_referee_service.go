package web_referee_service

import (
	"github.com/google/uuid"
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

type WebRefereeStorage interface {
	CreateUser(ctx context.Context, user *models.User) (uuid.UUID, error)
	CreateTournament(ctx context.Context, tournament *models.Tournament) (uuid.UUID, error)
	CreatePlayer(ctx context.Context, player *models.Player) (uuid.UUID, error)
	GetPlayersByTournamentId(ctx context.Context, tournamentId uuid.UUID) ([]*models.Player, error)
	CreatePairing(ctx context.Context, games []*models.Game, tournamentUUID uuid.UUID) error
	GetPairing(ctx context.Context, tournamentId uuid.UUID) ([]*models.Game, error)
}

type WebRefereeCache interface {
	SavePairing(ctx context.Context, tournamentID uuid.UUID, games []*models.Game) error
	GetPairing(ctx context.Context, tournamentID uuid.UUID) ([]*models.Game, error)
}

type WebRefereeProducer interface {
	ProduceUser(ctx context.Context, user models.User) error
}

type WebRefereeService struct {
	webRefereeStorage  WebRefereeStorage
	webRefereeProducer WebRefereeProducer
	webRefereeCache    WebRefereeCache
}

func NewWebRefereeService(webRefereeStorage WebRefereeStorage, producer WebRefereeProducer, cache WebRefereeCache) *WebRefereeService {
	return &WebRefereeService{
		webRefereeStorage:  webRefereeStorage,
		webRefereeProducer: producer,
		webRefereeCache:    cache,
	}
}
