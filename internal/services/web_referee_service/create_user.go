package web_referee_service

import (
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

func (service *WebRefereeService) CreateUser(ctx context.Context, user *models.User) error {
	err := service.webRefereeStorage.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
