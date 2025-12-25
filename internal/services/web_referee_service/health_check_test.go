package web_referee_service

import (
	"github.com/stretchr/testify/assert"
)

func (suite *WebRefereeServiceSuite) TestHealthCheck() {
	err := suite.webRefereeService.HealthCheck(suite.ctx)

	assert.Equal(suite.T(), nil, err)
}
