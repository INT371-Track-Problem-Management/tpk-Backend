FROM golang:1.17-alpine

WORKDIR ../app

COPY .. ./

RUN GOOS=linux go build app/cmd/tpkBackend.go

EXPOSE 5000

CMD KEY="DEV" ./tpkBackend
