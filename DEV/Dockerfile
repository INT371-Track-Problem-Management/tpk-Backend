FROM golang:1.17-alpine

WORKDIR ../app

RUN mkdir ../images
COPY .. ./

RUN GOOS=linux go build app/cmd/server.go

EXPOSE 3000

CMD KEY="DEV" ./server
