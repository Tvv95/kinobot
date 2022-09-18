FROM golang:latest

COPY . /kinobot/
WORKDIR /kinobot/

RUN go mod download
RUN go build -o ./.bin/bot cmd/app/main.go

CMD ["./.bin/bot"]