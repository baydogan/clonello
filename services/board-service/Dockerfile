FROM golang:1.23.3-alpine

WORKDIR /app

RUN apk add --no-cache protoc protobuf-dev

COPY go.mod go.sum ./

RUN go mod tidy

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY . .

RUN protoc \
    --go_out=./internal/proto \
    --go-grpc_out=./internal/proto \
    --proto_path=./internal/proto \
    ./internal/proto/board_service.proto

RUN go build -o board-service cmd/main.go

EXPOSE 8081

CMD ["./board-service"]
