package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Sound struct {
	Frequency int `json:"frequency"`
	Length    int `json:"length"`
}

func main() {
	time.Sleep(time.Second * 2)
	client := mqtt.NewClient(
		mqtt.NewClientOptions().
			AddBroker("tcp://mosquitto:1883").
			SetClientID("app_generator"),
	)

	if appToken := client.Connect(); appToken.Wait() && appToken.Error() != nil {
		panic(appToken.Error())
	}

	for i := 0; i < 100; i++ {
		time.Sleep(3 * time.Second)
		s := Sound{
			Frequency: rand.Intn(20) + 20,
			Length:    rand.Intn(3) + 1,
		}
		msg, _ := json.Marshal(s)
		token := client.Publish("iotvisual", 0, false, msg)
		fmt.Printf("Published message: %#v\n", s)
		fmt.Println("Everything's OK: ", token.Wait())
	}

}
