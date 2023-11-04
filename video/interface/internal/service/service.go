package service

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/video/internal/videoApi"
	"google.golang.org/grpc"
)

type Service struct {
	baseClient videoApi.VideoServiceClient
}

func NewService(conn *grpc.ClientConn) *Service {
	return &Service{baseClient: videoApi.NewVideoServiceClient(conn)}
}

func (s *Service) UploadVideoToken(ctx context.Context, req *videoApi.UploadVideoReq) (resp *videoApi.UploadVideoResp, err error) {
	return s.baseClient.UploadVideoToken(ctx, req)
}

func (s *Service) DownLoadVideoToken(ctx context.Context, req *videoApi.DownloadReq) (resp *videoApi.DownLoadResp, err error) {
	return s.baseClient.DownLoadVideoToken(ctx, req)
}
