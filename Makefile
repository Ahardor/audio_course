.PHONY:
	gen 

gen: 
	mkdir -p mock/internal
	protoc --go_out=mock/internal --go_opt=paths=source_relative \
		--go-grpc_out=mock/internal --go-grpc_opt=paths=source_relative \
		mock/api/mock_v1/service.proto