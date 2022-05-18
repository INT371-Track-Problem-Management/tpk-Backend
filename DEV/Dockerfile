FROM golang:1.17-alpine

WORKDIR ../app

COPY .. ./
COPY ../app/.env app/.env

RUN GOOS=linux go build app/cmd/tpkBackend.go

EXPOSE 3000

CMD KEY="DEV" ./tpkBackend
