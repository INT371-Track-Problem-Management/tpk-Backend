FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ../../ ./

EXPOSE 5000

CMD KEY="DEV" go build app/cmd/tpkBackend.go \
    && ./tpkBackend
