package server

import (
	"context"
	"fmt"
	"io"
	"iotvisual/processor/internal/cacher"
	"iotvisual/processor/internal/processor/api/processor_v1"
	"iotvisual/processor/internal/server/handlers"
	"log"
	"os"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/rs/zerolog"
)

type Server struct {
	Logger zerolog.Logger
	Mqtt   mqtt.Client
	Db     *mongo.Client
	cache  *cacher.MelodyCache
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
			log.Fatal("failed to initialize database: %w", err)
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

func WithCache() Option {
	return func(s *Server) {
		s.cache = cacher.New(
			cacher.WithExpirationTime(5*time.Minute),
			cacher.WithCleanupInterval(2*time.Minute),
		)
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
	client, err := mongo.Connect(
		ctx,
		options.Client().
			ApplyURI("mongodb://mongo").
			SetAuth(options.Credential{
				Username: "iotvisual",
				Password: "iotvisualpass",
			}),
	)

	if err != nil {
		return nil, err
	}
	col := client.Database("iotDB").Collection("Melodies")
	_, err = col.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.D{{Key: "id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)

	if err != nil {
		return nil, err
	}
	return client, nil
}

func initMQTT() (mqtt.Client, error) {
	opts := mqtt.NewClientOptions().
		SetDefaultPublishHandler(handlers.MelodyEventHandler()).
		AddBroker("tcp://mosquitto:1883").
		SetClientID("app_processor")
	client := mqtt.NewClient(opts)
	if appToken := client.Connect(); appToken.Wait() && appToken.Error() != nil {
		return nil, appToken.Error()
	}
	return client, nil
}
