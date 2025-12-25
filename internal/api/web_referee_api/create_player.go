package web_referee_service_api

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/tengenatari/web-referee/internal/models"
	"github.com/tengenatari/web-referee/internal/pb/web_referee_api"
	"golang.org/x/net/context"
)

func (s *WebRefereeServiceAPI) CreatePLayer(ctx context.Context, req *web_referee_api.CreatePlayerRequest) error {
	tournamentUUID, err := uuid.Parse(req.TournamentId)
	if err != nil {
		return errors.Wrap(err, "error parsing tournament id")
	}
	userUUID, err := uuid.Parse(req.UserId)
	if err != nil {
		return errors.Wrap(err, "error parsing user id")
	}

	player := &models.Player{
		TournamentId: tournamentUUID,
		UserId:       userUUID,
		MacMahon:     req.MacMahon,
	}

	err = s.webRefereeService.CreatePlayer(ctx, player)
	if err != nil {
		return errors.Wrap(err, "CreatePlayer failed")
	}
	return nil
}
