FROM golang:1.17-alpine

WORKDIR ../app

RUN mkdir ../images
RUN cd ../images && mkdir report
RUN cd ../images && mkdir profile
COPY .. ./

RUN GOOS=linux go build app/cmd/server.go

EXPOSE 3000

CMD KEY="DEV" ./server
