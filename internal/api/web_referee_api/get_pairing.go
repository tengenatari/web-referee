package web_referee_service_api

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/tengenatari/web-referee/internal/models"
	models2 "github.com/tengenatari/web-referee/internal/pb/models"
	"github.com/tengenatari/web-referee/internal/pb/web_referee_api"
	"golang.org/x/net/context"
)

func (s *WebRefereeServiceAPI) GetPairing(ctx context.Context, req *web_referee_api.GetPairingRequest) (*web_referee_api.GetPairingResponse, error) {

	tournamentUUID, err := uuid.Parse(req.TournamentId)

	if err != nil {
		return nil, errors.Wrap(err, "error parsing tournament uuid")
	}

	pairing, err := s.webRefereeService.GetPairing(ctx, tournamentUUID)
	if err != nil {
		return nil, errors.Wrap(err, "error creating pairing")
	}

	return &web_referee_api.GetPairingResponse{
		Games: lo.Map(pairing, func(game *models.Game, _ int) *models2.Game {
			return &models2.Game{
				Id:    game.Id.String(),
				White: game.White.String(),
				Black: game.Black.String(),
			}
		}),
	}, nil
}
