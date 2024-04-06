package main

import (
	"context"
	"fmt"
	"iotvisual/processor/internal/processor/api/processor_v1"
	"iotvisual/processor/internal/server"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	appCtx := context.Background()
	s := server.New(
		server.WithLogger(),
		//server.WithDatabase(),
		// server.WithMQTTClient(),
		server.WithCache(),
	)
	defer func() {
		if err := s.Db.Disconnect(appCtx); err != nil {
			panic(err)
		}
	}()

	// subToken := s.Mqtt.Subscribe("iotvisual", 1, handlers.MelodyEventHandler())
	// if ok := subToken.Wait(); !ok || subToken.Error() != nil {
	// 	s.Logger.Error().
	// 		Msgf("MQTT processor subscribe error: %s, token received: %t",
	// 			subToken.Error().Error(), ok,
	// 		)
	// }

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7259))
	if err != nil {
		s.Logger.Fatal().Msgf("server listen error: %s", err.Error())
	}
	serv := grpc.NewServer()
	reflection.Register(serv)
	processor_v1.RegisterProcessorServiceServer(serv, s)

	if err = serv.Serve(lis); err != nil {
		s.Logger.Fatal().Msgf("server start error: %s", err.Error())
	}
}
