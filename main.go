package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/paqstd-team/fake-cli/app"
)

func main() {
	Main()
}

// Main is the exported entrypoint to allow external tests to exercise startup logic.
func Main() {
	// Define command-line flags
	port := flag.Int("p", 8080, "port number")
	configPath := flag.String("c", "config.json", "path to config file")
	flag.Parse()

	// Check if PORT environment variable is set and use it if not overridden by a command-line flag
	envPort := os.Getenv("PORT")
	if envPort != "" {
		var err error
		*port, err = strconv.Atoi(envPort)
		if err != nil {
			log.Fatalf("Invalid port number: %v", envPort)
		}
	}

	// Check if CONFIG_PATH environment variable is set and use it if not overridden by a command-line flag
	envConfig := os.Getenv("CONFIG_PATH")
	if envConfig != "" {
		*configPath = envConfig
	}

	server, err := app.Run(*configPath, *port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	// Allow tests to execute main without blocking
	if os.Getenv("TESTING") == "1" {
		return
	}

	// Block in production run
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
