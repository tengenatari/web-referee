package web_referee_service

import (
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/tengenatari/web-referee/internal/models"
	"golang.org/x/net/context"
)

func (service *WebRefereeService) CreateUser(ctx context.Context, user *models.User) error {

	err := validateUser(user)
	if err != nil {
		return errors.Wrap(err, "validateUser failed")
	}

	err = service.webRefereeStorage.CreateUser(ctx, user)
	if err != nil {
		return errors.Wrap(err, "CreateUser failed")
	}

	err = service.webRefereeProducer.ProduceUser(ctx, *user)
	if err != nil {
		return errors.Wrap(err, "CreateUser failed")
	}

	return nil
}

func validateUser(user *models.User) error {
	if user == nil {
		return errors.New("user cannot be nil")
	}

	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email is required")
	}

	if !isValidEmail(user.Email) {
		return errors.New("invalid email format")
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

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
