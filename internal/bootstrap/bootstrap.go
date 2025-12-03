package bootstrap

import (
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tengenatari/web-referee/config"
	web_referee_service_api "github.com/tengenatari/web-referee/internal/api/web_referee_api"
	"github.com/tengenatari/web-referee/internal/pb/web_referee_api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AppRun(api *web_referee_service_api.WebRefereeServiceAPI, cfg *config.Config) {
	go func() {
		if err := runGRPCServer(api); err != nil {
			panic(fmt.Errorf("failed to run gRPC server: %v", err))
		}
	}()

	err := runGatewayServer(cfg.Web.Host, cfg.Web.Port)
	if err != nil {
		panic(fmt.Errorf("failed to run gateway server: %v", err))
	}

}

func runGRPCServer(api *web_referee_service_api.WebRefereeServiceAPI) error {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	web_referee_api.RegisterWebRefereeServiceServer(s, api)

	slog.Info("gRPC-server server listening on :50051")
	return s.Serve(lis)
}

func runGatewayServer(host string, port int) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	grpcAddress := fmt.Sprintf("%s:%d", host, 50051)
	httpAddress := fmt.Sprintf("%s:%d", host, port)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := web_referee_api.RegisterWebRefereeServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)

	if err != nil {
		return err
	}
	r := chi.NewRouter()
	/* swaggerPath := os.Getenv("swaggerPath")

	if _, err := os.Stat(swaggerPath); os.IsNotExist(err) {
		return fmt.Errorf("swagger file not found: %s", swaggerPath)
	}


	r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, swaggerPath)
	})

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger.json"),
	)) */

	r.Mount("/", mux)

	slog.Info(fmt.Sprintf("Listening on %s", httpAddress))
	return http.ListenAndServe(httpAddress, r)
}
