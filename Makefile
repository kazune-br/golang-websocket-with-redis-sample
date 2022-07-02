.PHONY: build up down logs

build: go-build
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker compose build

up: build
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f

lint:
	go fmt ./...
	golangci-lint run
	go vet ./...

go-build: lint
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./.artifacts/app-linux-amd64 ./cmd/app.go

open:
	open assets/index.html

publish:
	redis-cli --pass password -p 6380 publish sample "hello"
