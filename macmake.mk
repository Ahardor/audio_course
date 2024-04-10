.PHONY:	
	mac-up \
	mac-down \
	mac-purge \
	mac-test-processor

mac-up:
	docker-compose -f docker-compose.yaml up --build

mac-down:
	docker-compose -f docker-compose.yaml down

mac-purge:
	docker-compose down
	docker system prune -a

mac-test-processor:
	docker run -d -p 127.0.0.1:27017:27017 --name=mongotest \
		-e MONGO_INITDB_ROOT_USERNAME=iotvisual \
   		-e MONGO_INITDB_ROOT_PASSWORD=iotvisualpass \
		mongodb/mongodb-community-server
	go test ./processor/...
	docker stop mongotest && docker rm mongotest