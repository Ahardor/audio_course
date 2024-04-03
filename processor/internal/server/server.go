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
	Logger zerolog.Logger
	Mqtt   *mqtt.Client
	Db     *mongo.Client
	processor_v1.UnimplementedProcessorServiceServer
}

type Option func(s *Server)

func New(opts ...Option) *Server {
	s := &Server{}

	for i := range opts {
		opts[i](s)
	}

	return s
}

func WithLogger() Option {
	return func(s *Server) { s.Logger = initLogger(os.Stdout) }
}

func WithMQTTClient() Option {
	return func(s *Server) {
		client, err := initMQTT()
		if err != nil {
			fmt.Println("MQTT FUCKED UP")
			return
		}
		s.Mqtt = client
	}
}

func WithDatabase() Option {
	return func(s *Server) {
		client, err := initDatabase()
		if err != nil {
			log.Fatal("failed to initialize database: %w", err)
		}
		s.Db = client
	}
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

func initDatabase() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo"))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func initMQTT() (*mqtt.Client, error) {
	client := mqtt.NewClient(
		mqtt.NewClientOptions().
			AddBroker("tcp://mosquitto:1883").
			SetClientID("app_processor"),
	)
	if appToken := client.Connect(); appToken.Wait() && appToken.Error() != nil {
		return nil, appToken.Error()
	}
	return &client, nil
}
