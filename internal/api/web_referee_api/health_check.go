package web_referee_service_api

import (
	"github.com/tengenatari/web-referee/internal/pb/web_referee_api"
	"golang.org/x/net/context"
)

func (s *WebRefereeServiceAPI) HealthCheck(ctx context.Context, req *web_referee_api.HealthCheckRequest) (*web_referee_api.HealthCheckResponse, error) {
	err := s.webRefereeService.HealthCheck(ctx)
	if err != nil {
		return nil, err
	}
	return &web_referee_api.HealthCheckResponse{}, nil
}
