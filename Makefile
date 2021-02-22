run-server:
	realize start --server

run-client:
	cd client &&\
	yarn start

protoc-server:
	protoc -I protobuf/ protobuf/enums.proto --go_out=plugins=grpc:./
	protoc -I protobuf/ protobuf/messages.proto --go_out=plugins=grpc:./
	protoc -I protobuf/ protobuf/services.proto --go_out=plugins=grpc:./

docker-build:
	cd docker &&\
	docker-compose build --no-cache

docker-up:
	cd docker &&\
	docker-compose up -d

docker-stop:
	cd docker &&\
	docker-compose stop

docker-rm:
	cd docker &&\
	docker-compose rm -f
	rm -rf db/data/*

docker-ps:
	cd docker &&\
	docker-compose ps

docker-exec-db:
	cd docker &&\
	docker exec -it sns-db /bin/bash

docker-exec-proxy:
	cd docker &&\
	docker exec -it sns-proxy /bin/bash