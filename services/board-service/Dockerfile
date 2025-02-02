FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o board-service ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/board-service .

EXPOSE 50051

CMD ["./board-service"]
