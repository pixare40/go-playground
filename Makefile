.PHONY: protos

protos:
	protoc -I proto/ proto/*.proto --go_out=proto --go-grpc_out=proto --go-grpc_opt=require_unimplemented_servers=false