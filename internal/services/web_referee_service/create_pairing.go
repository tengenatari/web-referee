package web_referee_service

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

func (service *WebRefereeService) CreatePairing(ctx context.Context, tournamentId uuid.UUID) error {
	players, err := service.webRefereeStorage.GetPlayersByTournamentId(ctx, tournamentId)
	if err != nil {
		return errors.Wrap(err, "GetPlayersByTournamentId")
	}
	var pairings []*models.Game
	for i := 0; i < len(players)/2; i++ {
		pairings = append(pairings, &models.Game{
			White: players[i].Id,
			Black: players[i+len(players)/2].Id,
		})
	}
	err = service.webRefereeStorage.CreatePairing(ctx, pairings, tournamentId)
	if err != nil {
		return errors.Wrap(err, "CreatePairing")
	}
	return nil
}
