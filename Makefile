.PHONY:
	gen \
	up \
	down \
	purge \

gen: 
	mkdir -p mock/internal
	protoc --go_out=mock/internal --go_opt=paths=source_relative \
		--go-grpc_out=mock/internal --go-grpc_opt=paths=source_relative \
		mock/api/mock_v1/service.proto

up:
	sudo docker compose up

down:
	sudo docker compose down

purge:
	sudo docker compose down
	sudo docker system prune -a