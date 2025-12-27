package web_referee_service

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

func (service *WebRefereeService) CreateUser(ctx context.Context, user *models.User) (string, error) {

	err := validateUser(user)
	if err != nil {
		return "", errors.Wrap(err, "validateUser failed")
	}

	userUuid, err := service.webRefereeStorage.CreateUser(ctx, user)
	if err != nil {
		return "", errors.Wrap(err, "CreateUser failed")
	}
	user.Id = userUuid
	err = service.webRefereeProducer.ProduceUser(ctx, *user)
	if err != nil {
		return "", errors.Wrap(err, "CreateUser failed")
	}

	return userUuid.String(), nil
}

func validateUser(user *models.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	if strings.TrimSpace(user.Name) == "" {
		return errors.New("username is required")
	}

	if strings.TrimSpace(user.TigrId) == "" {
		return errors.New("tigr_id is required")
	}

	if user.Rating < 0 {
		return errors.New("rating must be greater than 0")
	}

	return nil
}
