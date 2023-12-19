generate_authorization:
	protoc --go_out=./gen/go/authorization/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/authorization --go-grpc_opt=paths=source_relative ./proto/authorization/authorization.proto

generate_user:
	protoc --go_out=./gen/go/user/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/user --go-grpc_opt=paths=source_relative ./proto/user/user.proto

generate: generate_authorization generate_user

push:
	git push --tags origin main