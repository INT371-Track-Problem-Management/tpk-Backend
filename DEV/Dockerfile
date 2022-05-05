FROM golang:1.17-alpine

COPY ../ ./
WORKDIR /app

COPY ../go.mod ./
COPY ../go.sum ./

CMD cd .. \ go mod download

EXPOSE 5000

CMD go build app/cmd/tpkBackend.go \
    && KEY="DEV"  ./tpkBackend
