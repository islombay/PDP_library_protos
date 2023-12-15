generate:
	protoc --go_out=./gen/go/authorization/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/authorization --go-grpc_opt=paths=source_relative ./proto/authorization/authorization.proto