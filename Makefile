check-env-file:
	@if [ -f .env ]; then \
		echo "The .env file is present."; \
	else \
		echo "Error: The .env file is missing."; \
		exit 1; \
	fi

build:
	go mod download && go mod verify
	go build -v -o bin/app

run: check-env-file build
	@./bin/app
