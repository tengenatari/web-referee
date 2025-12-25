package web_referee_service_api

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tengenatari/web-referee/mocks"
	"golang.org/x/net/context"
)

type WebRefereeServiceAPISuite struct {
	suite.Suite
	ctx                  context.Context
	webRefereeService    *mocks.WebRefereeService
	webRefereeServiceAPI *WebRefereeServiceAPI
}

func (suite *WebRefereeServiceAPISuite) SetupTest() {
	suite.webRefereeService = mocks.NewWebRefereeService(suite.T())
	suite.webRefereeServiceAPI = NewWebRefereeServiceAPI(suite.webRefereeService)
	suite.ctx = context.Background()
}

func TestRunAPISuite(t *testing.T) {
	suite.Run(t, new(WebRefereeServiceAPISuite))
}
