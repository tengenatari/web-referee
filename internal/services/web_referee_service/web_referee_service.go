package web_referee_service

import (
	"github.com/tengenatari/web-referee/internal/models"
	"github.com/tengenatari/web-referee/internal/storage/pgstorage"
	"golang.org/x/net/context"
)

type WebRefereeStorage interface {
	CreateUser(ctx context.Context, user *models.User) error
}

type WebRefereeService struct {
	webRefereeStorage WebRefereeStorage
}

func NewWebRefereeService(webRefereeStorage *pgstorage.PGstorage) *WebRefereeService {
	return &WebRefereeService{
		webRefereeStorage: webRefereeStorage,
	}
}
