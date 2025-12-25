package web_referee_service

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

func (service *WebRefereeService) CreateTournament(ctx context.Context, tournament *models.Tournament) error {
	err := validateTournament(ctx, tournament)
	if err != nil {
		return errors.Wrap(err, "error validating tournament")
	}
	err = service.webRefereeStorage.CreateTournament(ctx, tournament)
	if err != nil {
		return errors.Wrap(err, "error creating tournament")
	}
	return nil
}

func validateTournament(ctx context.Context, tournament *models.Tournament) error {
	if tournament == nil {
		return errors.New("Tournament is required")
	}
	if strings.TrimSpace(tournament.Name) == "" {
		return errors.New("Tournament name is required")
	}
	return nil
}
