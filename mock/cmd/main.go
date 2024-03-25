package main

import (
	"fmt"
	"iotvisual/mock/internal/mock/api/mock_v1"
	"iotvisual/mock/internal/server"
	"log"
	"net"

	//mqtt "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Sound struct {
	Frequency int `json:"frequency"`
	Length    int `json:"length"`
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 6969))
	if err != nil {
		log.Fatal(err.Error())
	}
	server := server.InitServer()
	s := grpc.NewServer()
	reflection.Register(s)
	mock_v1.RegisterMockServiceServer(s, server)

	if err = s.Serve(lis); err != nil {
		server.Logger.Fatal().Msgf("Server start error: %s", err.Error())
	}

	// time.Sleep(time.Second * 2)
	// client := mqtt.NewClient(
	// 	mqtt.NewClientOptions().
	// 		AddBroker("tcp://mosquitto:1883").
	// 		SetClientID("app_generator"),
	// )

	// if appToken := client.Connect(); appToken.Wait() && appToken.Error() != nil {
	// 	panic(appToken.Error())
	// }

	// for i := 0; i < 100; i++ {
	// 	time.Sleep(3 * time.Second)
	// 	s := Sound{
	// 		Frequency: rand.Intn(20) + 20,
	// 		Length:    rand.Intn(3) + 1,
	// 	}
	// 	msg, _ := json.Marshal(s)
	// 	token := client.Publish("iotvisual", 0, false, msg)
	// 	fmt.Printf("Published message: %#v\n", s)
	// 	fmt.Println("Everything's OK: ", token.Wait())
	// }

}
