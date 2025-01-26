FROM golang:1.23.3-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o board-service cmd/main.go

EXPOSE 8081

CMD ["./board-service"]
