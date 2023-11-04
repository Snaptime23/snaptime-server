package server

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/video/internal/videoApi"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/service"
)

type Server struct {
	svr *service.Service
	videoApi.UnimplementedVideoServiceServer
}

func NewServer() *Server {
	return &Server{
		svr:                             service.NewService(),
		UnimplementedVideoServiceServer: videoApi.UnimplementedVideoServiceServer{},
	}
}

func (s *Server) VideoFeed(ctx context.Context, req *videoApi.VideoFeedReq) (resp *videoApi.VideoFeedResp, err error) {
	return s.svr.VideoFeed(ctx, req)
}

func (s *Server) UploadVideoToken(ctx context.Context, req *videoApi.UploadVideoReq) (resp *videoApi.UploadVideoResp, err error) {
	return s.svr.UploadVideoToken(ctx, req)
}

func (s *Server) DownLoadVideoToken(ctx context.Context, req *videoApi.DownloadReq) (reso *videoApi.DownLoadResp, err error) {
	return s.svr.DownLoadVideoToken(ctx, req)
}
