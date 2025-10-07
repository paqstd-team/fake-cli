package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/paqstd-team/fake-cli/v2/config"
	"github.com/paqstd-team/fake-cli/v2/handler"
)

// Run constructs the HTTP server using the provided config path and port.
// It seeds the random generator to make responses deterministic across runs.
func Run(configPath string, port int) (*http.Server, error) {
	gofakeit.Seed(0)

	cfg, err := config.LoadConfigFromFile(configPath)
	if err != nil {
		return nil, err
	}

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: handler.MakeHandler(cfg),
	}

	log.Printf("Starting server on %v", srv.Addr)
	return srv, nil
}
