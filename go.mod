module iotvisual

go 1.21.7

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.13.0

require (
	github.com/eclipse/paho.mqtt.golang v1.4.3
	github.com/rs/zerolog v1.32.0
	google.golang.org/grpc v1.62.1
	google.golang.org/protobuf v1.33.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240325164216-beb30f47624b // indirect
)
