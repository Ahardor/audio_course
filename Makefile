include macmake.mk

.PHONY:
	gen-mock \
	gen-processor \
	up \
	down \
	purge \
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
	sudo docker compose -f docker-compose.yaml up --build

down:
	sudo docker compose -f docker-compose.yaml down

purge:
	sudo docker compose down
	sudo docker system prune -a