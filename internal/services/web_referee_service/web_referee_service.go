package web_referee_service

type WebRefereeStorage interface {
}

type WebRefereeService struct {
	webRefereeStorage WebRefereeStorage
}

func NewWebRefereeService(webRefereeStorage WebRefereeStorage) *WebRefereeService {
	return &WebRefereeService{
		webRefereeStorage: webRefereeStorage,
	}
}
