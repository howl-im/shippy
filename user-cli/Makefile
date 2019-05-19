all: build run

build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t user-cli .

run:
	docker run -e MICRO_REGISTRY=mdns user-cli command \
		--name="Evan Liu" \
		--email="evan@bbc.com" \
		--password="testing123" \
		--company="BBC"

.PHONY: all build run
