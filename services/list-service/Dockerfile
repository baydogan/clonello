FROM golang:1.23.3-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o list-service cmd/main.go

EXPOSE 50052

CMD ["./list-service"]