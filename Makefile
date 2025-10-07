# Define the name of the binary
BINARY_NAME=fake-cli

# Define the path to the main source file
MAIN_PATH=./main.go

# Define the command-line flags to pass to the build command
BUILD_FLAGS=-ldflags="-s -w"

# Define the build command
build:
	go build $(BUILD_FLAGS) -o $(BINARY_NAME) $(MAIN_PATH)

# Define the clean command
clean:
	rm -f $(BINARY_NAME)

# Define the run command
run:
	./$(BINARY_NAME)

# Define the default command as the build command
default: build

test:
	TESTING=1 go test ./... -covermode=atomic -coverpkg=./app,./cache,./config,./handler -coverprofile=coverage.out
	@echo "\nCoverage by function/package:" && go tool cover -func=coverage.out | sed 's/^/  /'
	@echo "\nEnforcing 100% coverage"
	@go tool cover -func=coverage.out | awk '/total:/ { if ($$3 != "100.0%") { print "ERROR: Coverage is not 100%"; exit 1 } }'
