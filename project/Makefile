API_BINARY=apiServer.out
EMAIL_BINARY=emailService.out

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_api build_email
	@echo "Stopping docker images (if running...)"
	docker-compose -f docker-compose.development.yaml down
	@echo "Building (when required) and starting docker images..."
	docker-compose -f docker-compose.development.yaml up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose -f docker-compose.development.yaml down
	@echo "Done!"

## build_api: builds the api binary as a linux executable
build_api:
	@echo "Building api binary..."
	cd ../api-service && env GOOS=linux CGO_ENABLED=0 go build -o ${API_BINARY} ./cmd
	@echo "Done!"

## build_email: builds the email binary as a linux executable
build_email:
	@echo "Building email binary..."
	cd ../email-service && env GOOS=linux CGO_ENABLED=0 go build -o ${EMAIL_BINARY} ./cmd
	@echo "Done!"