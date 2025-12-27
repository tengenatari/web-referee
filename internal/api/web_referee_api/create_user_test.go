package web_referee_service_api

import (
	"github.com/stretchr/testify/assert"
	"github.com/tengenatari/web-referee/internal/models"
	"github.com/tengenatari/web-referee/internal/pb/web_referee_api"
)

func (suite *WebRefereeServiceAPISuite) TestCreateUser() {

	req := &web_referee_api.CreateUserRequest{}

	suite.webRefereeService.On("CreateUser", suite.ctx, &models.User{}).
		Return(nil).
		Once()
	_, err := suite.webRefereeServiceAPI.CreateUser(suite.ctx, req)

	assert.NoError(suite.T(), err)
}
