package server

import (
	"context"
	"fmt"
	"iotvisual/mock/internal/mock/api/mock_v1"
	"os"

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
	fmt.Println(notes)
	return nil, nil
}
