# Define the name of the binary
BINARY_NAME=fake

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
