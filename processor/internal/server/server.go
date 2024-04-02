package server

import (
	"context"
	"fmt"
	"io"
	"iotvisual/processor/internal/processor/api/processor_v1"
	"log"
	"os"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/rs/zerolog"
)

type Server struct {
	Logger     zerolog.Logger
	MqttClient mqtt.Client
	client     *mongo.Client
	processor_v1.UnimplementedProcessorServiceServer
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
			SetClientID("app_processor"),
	)
	if appToken := server.MqttClient.Connect(); appToken.Wait() && appToken.Error() != nil {
		panic(appToken.Error())
	}
	server.initDatabase()
	return &server
}

func (s *Server) initDatabase() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo"))
	if err != nil {
		log.Fatal(err)
	}
	s.client = client
}
