FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal

RUN go build -o rest-server ./cmd/main.go
EXPOSE 8080
FROM gcr.io/distroless/base

COPY --from=builder /app/rest-server /usr/local/bin/rest-server

ENTRYPOINT ["/usr/local/bin/rest-server"]