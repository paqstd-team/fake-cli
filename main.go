package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/paqstd-team/fake-cli/config"
	"github.com/paqstd-team/fake-cli/handler"
)

func main() {
	// Seed the random number generator for reproducibility
	gofakeit.Seed(0)

	// Define command-line flags
	port := flag.Int("p", 8080, "port number")
	configPath := flag.String("c", "config.fake", "path to config file")
	flag.Parse()

	config, err := config.LoadConfigFromFile(*configPath)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(*port),
		Handler: handler.MakeHandler(config),
	}

	log.Printf("Starting server on %v", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
