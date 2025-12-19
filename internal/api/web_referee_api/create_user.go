package web_referee_service_api

import (
	"github.com/tengenatari/web-referee/internal/models"
	"github.com/tengenatari/web-referee/internal/pb/web_referee_api"
	"golang.org/x/net/context"
)

func (s *WebRefereeServiceAPI) CreateUser(ctx context.Context, req *web_referee_api.CreateUserRequest) (*web_referee_api.CreateUserResponse, error) {

	user := models.User{
		Rating: req.GetRating(),
		Name:   req.GetName(),
		TigrId: req.GetTigrId(),
	}
	err := s.webRefereeService.CreateUser(ctx, &user)
	if err != nil {
		return nil, err
	}
	return &web_referee_api.CreateUserResponse{}, nil
}
