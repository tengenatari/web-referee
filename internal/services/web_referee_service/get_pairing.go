package web_referee_service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

func (service *WebRefereeService) GetPairing(ctx context.Context, tournamentId uuid.UUID) ([]*models.Game, error) {

	getPairing, err := service.webRefereeCache.GetPairing(ctx, tournamentId)
	if err == nil {
		return getPairing, nil
	}
	fmt.Println("GetPairing err: ", err)
	pairing, err := service.webRefereeStorage.GetPairing(ctx, tournamentId)
	if err != nil {
		return nil, errors.Wrap(err, "GetPairing")
	}

	err = service.webRefereeCache.SavePairing(ctx, tournamentId, pairing)
	if err != nil {
		fmt.Printf("Error caching pairing: %v\n", err)
		err = nil
	}

	return pairing, err
}
