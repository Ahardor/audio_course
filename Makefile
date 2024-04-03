.PHONY:
	gen-mock \
	gen-processor \
	up \
	mac-up \
	down \
	mac-down \
	purge \
	mac-purge \
	local-db-up \
	local-db-down

gen-mock: 
	mkdir -p mock/internal
	protoc --go_out=mock/internal --go_opt=paths=source_relative \
		--go-grpc_out=mock/internal --go-grpc_opt=paths=source_relative \
		mock/api/mock_v1/service.proto

gen-processor: 
	mkdir -p processor/internal
	protoc --go_out=processor/internal --go_opt=paths=source_relative \
		--go-grpc_out=processor/internal --go-grpc_opt=paths=source_relative \
		processor/api/processor_v1/service.proto

up:
	sudo docker compose up

down:
	sudo docker compose down

purge:
	sudo docker compose down
	sudo docker system prune -a

mac-up:
	docker-compose up

mac-down:
	docker-compose down

mac-purge:
	docker-compose down
	docker system prune -a

local-db-up:
	sudo docker run -d -p 127.0.0.1:27017:27017 --name=mongotest  mongodb/mongodb-community-server

local-db-down:
	sudo docker stop mongotest && sudo docker rm mongotest