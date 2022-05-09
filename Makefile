all: setup compile build up

setup:
	go mod tidy

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main cmd/api/main.go

build:
	docker build . -t golang-project-a:latest

up:
	docker-compose up


