.DEFAULT_GOAL := run

run :
	@go run main.go

update_proto :
	@protoc ./plugin/shared/grpc/plugin.proto --go-grpc_out=. --go_out=. --go-grpc_opt=require_unimplemented_servers=false

.PHONY: run