FROM golang:1.21.9-alpine3.18 AS builder

COPY . /github.com/kirillmc/GRPC_TEST_SERVER/source/
WORKDIR /github.com/kirillmc/GRPC_TEST_SERVER/source/

RUN go mod download
RUN go build -o ./bin/GRPC_TEST_SERVER cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/kirillmc/GRPC_TEST_SERVER/source/bin/GRPC_TEST_SERVER .
COPY .env .
CMD ["./GRPC_TEST_SERVER"]



