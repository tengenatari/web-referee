package web_referee_service_api

import (
	"time"

	"github.com/pkg/errors"
	"github.com/tengenatari/web-referee/internal/models"
	"github.com/tengenatari/web-referee/internal/pb/web_referee_api"
	"golang.org/x/net/context"
)

func (s *WebRefereeServiceAPI) CreateTournament(ctx context.Context, req *web_referee_api.CreateTournamentRequest) (*web_referee_api.CreateTournamentResponse, error) {

	parse, err := time.Parse(time.RFC3339, req.Date)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse date")
	}
	tournament := &models.Tournament{
		Name: req.Name,
		Date: parse,
	}

	tournamentUUID, err := s.webRefereeService.CreateTournament(ctx, tournament)
	if err != nil {
		return nil, errors.Wrap(err, "CreateTournament failed")
	}

	return &web_referee_api.CreateTournamentResponse{TournamentUUID: tournamentUUID}, nil
}
