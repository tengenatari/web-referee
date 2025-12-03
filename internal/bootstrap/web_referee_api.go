package bootstrap

import (
	webRefereeServiceApi "github.com/tengenatari/web-referee/internal/api/web_referee_api"
	"github.com/tengenatari/web-referee/internal/services/web_referee_service"
)

func InitWebRefereeServiceAPI(webRefereeService *web_referee_service.WebRefereeService) *webRefereeServiceApi.WebRefereeServiceAPI {
	return webRefereeServiceApi.NewWebRefereeServiceAPI(webRefereeService)
}
