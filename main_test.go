package main_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/paqstd-team/fake-cli/config"
	"github.com/paqstd-team/fake-cli/handler"
)

func TestMain(m *testing.M) {
	// Create temporary config file for testing
	configData := []byte(`{
		"cache": 10,
		"endpoints": [
			{ "url": "/test", "fields": { "foo": "sentence" } }
		]
	}`)
	configFile, err := ioutil.TempFile("", "config.*.json")
	if err != nil {
		panic(err)
	}
	defer os.Remove(configFile.Name())
	if _, err := configFile.Write(configData); err != nil {
		panic(err)
	}
	if err := configFile.Close(); err != nil {
		panic(err)
	}

	config, err := config.LoadConfigFromFile(configFile.Name())
	if err != nil {
		panic("Error loading config")
	}

	// Start the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler.MakeHandler(config),
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	defer server.Shutdown(nil)

	// Run the tests
	exitCode := m.Run()

	// Exit
	os.Exit(exitCode)
}

func TestServer(t *testing.T) {
	// Send a request to the server and check the response
	resp, err := http.Get("http://localhost:8080/test")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(body, []byte(`"foo"`)) {
		t.Errorf("expected response to contain \"foo\", got %s", body)
	}
}
