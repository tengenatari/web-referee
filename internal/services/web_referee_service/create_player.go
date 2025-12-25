package web_referee_service

import (
	"github.com/pkg/errors"
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

func (service *WebRefereeService) CreatePlayer(ctx context.Context, player *models.Player) error {
	err := service.webRefereeStorage.CreatePlayer(ctx, player)
	if err != nil {
		return errors.Wrap(err, "failed to create player")
	}
	return nil
}
