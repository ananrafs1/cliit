.DEFAULT_GOAL := run

bin := cli

build:
	@go build -o ./bin/cliit ./pkg/${bin}/main.go 

run : build
	@./bin/cliit

update_proto :
	@protoc ./plugin/shared/grpc/plugin.proto --go-grpc_out=. --go_out=. --go-grpc_opt=require_unimplemented_servers=false

.PHONY: run, build