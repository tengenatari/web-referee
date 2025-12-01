package bootstrap

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/tengenatari/web-referee/internal/pb/web_referee_api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AppRun() {
	err := runGatewayServer()
}

func runGatewayServer(host string, port string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	address := fmt.Sprintf("%s:%s", host, port)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := web_referee_api.RegisterWebRefereeServiceHandlerFromEndpoint(ctx, mux, address, opts)

	if err != nil {
		return err
	}

	swaggerPath := os.Getenv("swaggerPath")

	if _, err := os.Stat(swaggerPath); os.IsNotExist(err) {
		return fmt.Errorf("swagger file not found: %s", swaggerPath)
	}

	r := chi.NewRouter()
	r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, swaggerPath)
	})

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger.json"),
	))

	r.Mount("/", mux)

	slog.Info(fmt.Sprintf("Listening on %s", address))
	return http.ListenAndServe(address, r)
}
