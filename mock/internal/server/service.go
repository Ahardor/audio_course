package server

import (
	"context"
	"encoding/json"
	"fmt"
	"iotvisual/mock/internal/domain/messages"
	"iotvisual/mock/internal/mock/api/mock_v1"
	"os"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	"gopkg.in/yaml.v3"
)

func (s *Server) GetSoundFile(ctx context.Context, request *mock_v1.GetSoundFileRequest) (*emptypb.Empty, error) {
	if request == nil {
		return nil, fmt.Errorf("no message in request: GetSoundFile")
	}

	bytes, err := os.ReadFile(request.GetFilePath())
	if err != nil {
		return nil, err
	}

	var inputs []messages.MessageSoundInput
	if err = yaml.Unmarshal(bytes, &inputs); err != nil {
		return nil, err
	}

	for _, input := range inputs {
		msg, err := json.Marshal(input)
		if err != nil {
			s.Logger.Err(err).Msg("Error on json marshall")
			continue
		}
		token := s.MqttClient.Publish("sound/note", 1, false, msg)
		s.Logger.Debug().Msgf("Sending result: %t, with error: %v", token.Wait(), token.Error())
		time.Sleep(time.Duration(input.LengthMS) * time.Millisecond)
	}

	return &emptypb.Empty{}, nil
}
