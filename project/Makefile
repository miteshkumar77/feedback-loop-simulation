protos:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pxguide/px_guide.proto

srv:
	go build -o bin/server ./server/main.go

cli:
	go build -o bin/client ./client/main.go

clean:
	rm -rf bin/*

run_build:
	./bin/main