package web_referee_service_api

import (
	"github.com/stretchr/testify/assert"
	"github.com/tengenatari/web-referee/internal/pb/web_referee_api"
)

func (suite *WebRefereeServiceAPISuite) TestHealthCheck() {

	req := &web_referee_api.HealthCheckRequest{}

	suite.webRefereeService.On("HealthCheck", suite.ctx).
		Return(nil).
		Once()
	check, err := suite.webRefereeServiceAPI.HealthCheck(suite.ctx, req)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), &web_referee_api.HealthCheckResponse{}, check)
}
