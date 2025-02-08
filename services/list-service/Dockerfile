FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY services/list-service/go.mod services/list-service/go.sum ./

RUN go mod tidy && go mod download

COPY proto /proto

COPY services/list-service /app

RUN go build -o list-service ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/list-service .

EXPOSE 50052

CMD ["./list-service"]
