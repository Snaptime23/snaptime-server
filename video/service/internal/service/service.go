package service

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/video/internal/videoApi"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) VideoFeed(ctx context.Context, req *videoApi.VideoFeedReq) (resp *videoApi.VideoFeedResp, err error) {
	return
}
