package server

import (
	"context"
	"iotvisual/processor/internal/processor/api/processor_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) GetMockTemplate(ctx context.Context, r *processor_v1.GetMockTemplateRequest) (*emptypb.Empty, error) {
	return nil, nil
}
