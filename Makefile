run-server:
	cd server &&\
	go run main.go

run-client:
	cd client &&\
	yarn start

protoc-server:
	protoc -I protobuf/ protobuf/enums.proto --go_out=plugins=grpc:./
	protoc -I protobuf/ protobuf/messages.proto --go_out=plugins=grpc:./
	protoc -I protobuf/ protobuf/services.proto --go_out=plugins=grpc:./