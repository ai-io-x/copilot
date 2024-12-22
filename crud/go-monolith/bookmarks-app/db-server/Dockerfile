FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o db-server ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/db-server .

CMD ["./db-server"]