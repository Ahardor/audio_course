package server

import (
	"context"
	"encoding/json"
	"fmt"
	"iotvisual/mock/internal/mock/api/mock_v1"
	"os"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/yaml.v3"
)

type Note struct {
	Frequency int `yaml:"frequency"`
	Length    int `yaml:"length"`
}

func (s *Server) GetSoundFile(ctx context.Context, request *mock_v1.GetSoundFileRequest) (*emptypb.Empty, error) {
	if request == nil {
		return nil, fmt.Errorf("no message in request: GetSoundFile")
	}

	bytes, err := os.ReadFile(request.GetFilePath())
	if err != nil {
		return nil, err
	}

	var notes []Note
	if err = yaml.Unmarshal(bytes, &notes); err != nil {
		return nil, err
	}

	for _, note := range notes {
		msg, err := json.Marshal(note)
		if err != nil {
			s.Logger.Err(err).Msg("Error on json marshall")
			continue
		}
		token := s.MqttClient.Publish("iotvisual", 0, false, msg)
		s.Logger.Info().Msgf("Sending result: %t, with error: %v", token.Wait(), token.Error())
		time.Sleep(time.Duration(note.Length) * time.Second)
	}

	return nil, nil
}
