FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY api-gateway/go.mod api-gateway/go.sum ./

RUN go mod tidy && go mod download

COPY proto /proto

COPY api-gateway /app

RUN go build -o api-gateway ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/api-gateway .

EXPOSE 8082

CMD ["./api-gateway"]
