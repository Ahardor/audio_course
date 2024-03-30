package main

import (
	"fmt"
	"iotvisual/processor/internal/processor/api/processor_v1"
	"iotvisual/processor/internal/server"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	server := server.InitServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7259))
	if err != nil {
		log.Fatal(err.Error())
	}
	s := grpc.NewServer()
	reflection.Register(s)
	processor_v1.RegisterProcessorServiceServer(s, server)

	if err = s.Serve(lis); err != nil {
		server.Logger.Fatal().Msgf("Server start error: %s", err.Error())
	}
}
