.PHONY:	
	mup \
	mdown \
	mpurge \
	mtest-processor

mup:
	docker-compose -f docker-compose.yaml up --build

mdown:
	docker-compose -f docker-compose.yaml down

mpurge:
	docker-compose down
	docker system prune -a

mtest-processor:
	docker run -d -p 127.0.0.1:27017:27017 --name=mongotest \
		-e MONGO_INITDB_ROOT_USERNAME=iotvisual \
   		-e MONGO_INITDB_ROOT_PASSWORD=iotvisualpass \
		mongodb/mongodb-community-server
	go test ./processor/...
	docker stop mongotest && docker rm mongotest