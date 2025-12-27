package web_referee_service_api

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/tengenatari/web-referee/internal/pb/web_referee_api"
	"golang.org/x/net/context"
)

func (s *WebRefereeServiceAPI) CreatePairing(ctx context.Context, req *web_referee_api.CreatePairingRequest) (*web_referee_api.CreatePairingResponse, error) {

	tournamentUUID, err := uuid.Parse(req.TournamentId)

	if err != nil {
		return nil, errors.Wrap(err, "error parsing tournament uuid")
	}

	err = s.webRefereeService.CreatePairing(ctx, tournamentUUID)
	if err != nil {
		return nil, errors.Wrap(err, "error creating pairing")
	}
	return &web_referee_api.CreatePairingResponse{}, nil
}
