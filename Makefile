build:
	go build -o ./.bin/bot cmd/app/main.go

run: build
	./.bin/bot

build-image:
	docker build -t kinobot:v1.0 .

start-container:
	docker run kinobot:v1.0