package web_referee_service

import (
	"github.com/stretchr/testify/assert"
	"github.com/tengenatari/web-referee/internal/models"
)

func (suite *WebRefereeServiceSuite) TestCreateValidUser() {
	user := models.User{
		Email:  "vasya@mail.ru",
		Name:   "Andrey",
		TigrId: "125",
		Rating: 125,
	}

	suite.webRefereeStorage.On("CreateUser", suite.ctx, &user).
		Return(nil).
		Once()

	suite.webRefereeProducer.On("ProduceUser", suite.ctx, user).
		Return(nil).
		Once()

	err := suite.webRefereeService.CreateUser(suite.ctx, &user)

	assert.Nil(suite.T(), err)
}

func (suite *WebRefereeServiceSuite) TestCreateUserWithoutEmail() {
	user := models.User{
		Email:  "",
		Name:   "Andrey",
		TigrId: "125",
		Rating: 125,
	}
	wantErrorMessage := "validateUser failed: email is required"

	err := suite.webRefereeService.CreateUser(suite.ctx, &user)

	assert.EqualError(suite.T(), err, wantErrorMessage)
}

func (suite *WebRefereeServiceSuite) TestCreateUserWithIncorrectEmail() {
	user := models.User{
		Email:  "vasilmail.ru",
		Name:   "Andrey",
		TigrId: "125",
		Rating: 125,
	}
	wantErrorMessage := "validateUser failed: invalid email format"

	err := suite.webRefereeService.CreateUser(suite.ctx, &user)

	assert.EqualError(suite.T(), err, wantErrorMessage)
}

func (suite *WebRefereeServiceSuite) TestCreateUserWithIncorrectRating() {
	user := models.User{
		Email:  "vasil@mail.ru",
		Name:   "Andrey",
		TigrId: "125",
		Rating: -1,
	}
	wantErrorMessage := "validateUser failed: rating must be greater than 0"

	err := suite.webRefereeService.CreateUser(suite.ctx, &user)

	assert.EqualError(suite.T(), err, wantErrorMessage)
}

func (suite *WebRefereeServiceSuite) TestCreateUserWithoutTigrId() {
	user := models.User{
		Email:  "vasil@mail.ru",
		Name:   "Andrey",
		Rating: 100,
	}
	wantErrorMessage := "validateUser failed: tigr_id is required"

	err := suite.webRefereeService.CreateUser(suite.ctx, &user)

	assert.EqualError(suite.T(), err, wantErrorMessage)
}
