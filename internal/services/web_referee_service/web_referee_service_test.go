package web_referee_service

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/tengenatari/web-referee/mocks"
	"golang.org/x/net/context"
)

type WebRefereeServiceSuite struct {
	suite.Suite
	ctx                context.Context
	webRefereeService  *WebRefereeService
	webRefereeProducer *mocks.WebRefereeProducer
	webRefereeStorage  *mocks.WebRefereeStorage
}

func (suite *WebRefereeServiceSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.webRefereeStorage = mocks.NewWebRefereeStorage(suite.T())
	suite.webRefereeProducer = mocks.NewWebRefereeProducer(suite.T())
	suite.webRefereeService = NewWebRefereeService(suite.webRefereeStorage, suite.webRefereeProducer)
}

func TestRunServiceSuite(t *testing.T) {
	suite.Run(t, new(WebRefereeServiceSuite))
}
