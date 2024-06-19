package server

import (
	"fmt"
	"io"
	"iotvisual/mock/internal/mock/api/mock_v1"
	"os"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/rs/zerolog"
)

type Server struct {
	Logger     zerolog.Logger
	MqttClient mqtt.Client
	mock_v1.UnimplementedMockServiceServer
}

func initLogger(src io.Writer) zerolog.Logger {
	logger := zerolog.New(zerolog.ConsoleWriter{
		Out:        src,
		NoColor:    false,
		TimeFormat: time.ANSIC,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
		FormatTimestamp: func(i interface{}) string {
			t, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", i))
			return t.Format(time.RFC1123)
		},
	}).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	return logger
}

func InitServer() *Server {
	server := Server{}
	server.Logger = initLogger(os.Stdout)

	server.MqttClient = mqtt.NewClient(
		mqtt.NewClientOptions().
			AddBroker("tcp://mosquitto:1883").
			SetClientID("app_generator"),
	)

	if appToken := server.MqttClient.Connect(); appToken.Wait() && appToken.Error() != nil {
		panic(appToken.Error())
	}

	return &server
}
