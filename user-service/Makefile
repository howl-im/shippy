all: build run

build:
	@echo "Building~~~~~~~~~~~~~~~~"

	protoc -I . --go_out=plugins=micro:$(GOPATH)/src/shippy/user-service proto/user/user.proto

	GOOS=linux GOARCH=amd64 go build

	docker build -t user-service .

run:
	@echo "Runing~~~~~~~~~~~~~~~~~~~"

	docker run -p 50053:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=mdns \
		user-service

.PHONY: all build run
