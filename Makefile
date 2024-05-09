LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-program-api

generate-program-api:
	mkdir -p pkg/program_v3
	protoc --proto_path api/program_v3 \
	--go_out=pkg/program_v3 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/program_v3 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/program_v3/program_v3.proto

docker-build:
	docker buildx build --no-cache --platform linux/amd64 -t grpc_test_server:v0.0.1 .

grpc-load-test-get:
	ghz \
		--proto api/program_v3/program_v3.proto \
		--call program_v3.ProgramV3.Get \
		--data '{"Count": 5}' \
		--concurrency 10 \
		--total 3000 \
		--insecure \
		localhost:50051
grpc-load-test-get:
	ghz \
		--proto api/program_v3/program_v3.proto \
		--call program_v3.ProgramV3.Get \
		--data '{"Count": 5}' \
		--concurrency 10 \
		--total 3000 \
		--insecure \
		localhost:50051
grpc-load-test-post:
	ghz \
		--proto api/program_v3/program_v3.proto \
		--call program_v3.ProgramV3.Get \
		--data '{"Count": 5}' \
		--concurrency 10 \
		--total 3000 \
		--insecure \
		localhost:50051
grpc-load-test-delete:
	ghz \
		--proto api/program_v3/program_v3.proto \
		--call program_v3.ProgramV3.Get \
		--data '{"Count": 5}' \
		--concurrency 10 \
		--total 3000 \
		--insecure \
		localhost:50051

grpc-load-test1:
	ghz \
        --insecure \
        --proto api/program_v3/program_v3.proto \
        --call program_v3.ProgramV3.Get \
        --data '{"Count": 21}' \
        --total 1000 \
        --concurrency 100 \
        --connections 100 \
        --timeout 0 \
       localhost:50051

grpc-error-load-test:
	ghz \
		--proto api/program_v3/program_v3.proto \
		--call program_v3.ProgramV1.Get \
		--data '{"Count": 0}' \
		--rps 100 \
		--total 3000 \
		--insecure \
		localhost:50051

