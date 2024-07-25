LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-note-api

generate-note-api:
	mkdir -p pkg/chat_v1
	protoc --proto_path proto/chat/ \
	--go_out=pkg/chat_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chat_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	proto/chat/chat_v1.proto

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml


build-linux:
	GOOS=linux GOARCH=amd64 go build -o chat_serever_linux cmd/chat_server/app.go

copy-to-server:
	scp chat_serever_linux root@185.91.54.245:

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/courses/chat-server:v0.0.1 .
	docker login -u token -p token cr.selcloud.ru/courses
	docker push cr.selcloud.ru/courses/chat-server:v0.0.1
