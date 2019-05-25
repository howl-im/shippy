all: build run

build:
	@echo "Building~~~~~~~~~~~~~~~~"

	GOOS=linux GOARCH=amd64 go build

	docker build -t email-service .

run:
	@echo "Runing~~~~~~~~~~~~~~~~~~~"

	docker run -p 50054:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		-e MICRO_BROKER_ADDRESS=0.0.0.0:4222 \
		email-service

.PHONY: all build run
