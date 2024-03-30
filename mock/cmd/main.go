package main

import (
	"fmt"
	"iotvisual/mock/internal/mock/api/mock_v1"
	"iotvisual/mock/internal/server"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	server := server.InitServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7249))
	if err != nil {
		log.Fatal(err.Error())
	}
	s := grpc.NewServer()
	reflection.Register(s)
	mock_v1.RegisterMockServiceServer(s, server)

	if err = s.Serve(lis); err != nil {
		server.Logger.Fatal().Msgf("Server start error: %s", err.Error())
	}
}
