package service

import (
	"github.com/Snaptime23/snaptime-server/v2/video/internal/videoApi"
	"google.golang.org/grpc"
)

type Service struct {
	baseClient videoApi.VideoServiceClient
}

func NewService(conn *grpc.ClientConn) *Service {
	return &Service{baseClient: videoApi.NewVideoServiceClient(conn)}
}
