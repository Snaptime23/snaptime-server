package service

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/video/internal/videoApi"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/service/downloadToken"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/service/uploadToken"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) VideoFeed(ctx context.Context, req *videoApi.VideoFeedReq) (resp *videoApi.VideoFeedResp, err error) {
	return
}

func (s *Service) UploadVideoToken(ctx context.Context, req *videoApi.UploadVideoReq) (resp *videoApi.UploadVideoResp, err error) {
	resp = new(videoApi.UploadVideoResp)
	resp.Token = downloadToken.GetToken()
	return
}

func (s *Service) DownLoadVideoToken(ctx context.Context, req *videoApi.DownloadReq) (resp *videoApi.DownLoadResp, err error) {
	resp = new(videoApi.DownLoadResp)
	resp.Token = uploadToken.GetToken()
	return
}
