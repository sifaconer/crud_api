FROM golang:alpine AS build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . ./

RUN go build -o /api ./cmd/main.go

EXPOSE 8998

ENTRYPOINT ["/api"]