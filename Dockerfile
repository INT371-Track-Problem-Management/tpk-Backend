FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

EXPOSE 5000

CMD go build app/cmd/tpkBackend.go \
    && KEY="DEV"  ./tpkBackend
