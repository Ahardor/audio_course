package handlers

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func MelodyEventHandler() mqtt.MessageHandler {
	return func(c mqtt.Client, m mqtt.Message) {
		if !c.IsConnected() {
			return
		}
		fmt.Println(m)
	}
}
