FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o list-service ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/list-service .

EXPOSE 50052

CMD ["./list-service"]
