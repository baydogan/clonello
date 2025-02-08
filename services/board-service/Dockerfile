FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY services/board-service/go.mod services/board-service/go.sum ./

RUN go mod tidy && go mod download

COPY proto /proto

COPY services/board-service /app

RUN go build -o board-service ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/board-service .

EXPOSE 50051

CMD ["./board-service"]
