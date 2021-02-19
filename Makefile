run-server:
	cd server &&\
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

docker-init-db:
	docker-compose exec db /bin/bash -psecret -c "chmod 0775 docker-entrypoint-initdb.d/init-db.sh" &&\
	docker-compose exec db /bin/bash -psecret -c "sh ./docker-entrypoint-initdb.d/init-db.sh"

docker-reset-dummy:
	docker-compose exec db /bin/bash -psecret -c "chmod 0775 docker-entrypoint-initdb.d/reset_dummy.sh" &&\
	docker-compose exec db /bin/bash -psecret -c "sh ./docker-entrypoint-initdb.d/reset_dummy.sh"

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