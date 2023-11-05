package server

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/video/rpc_pb/videoApi"
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

func (s *Server) UploadVideo(ctx context.Context, req *videoApi.UploadVideoReq) (resp *videoApi.UploadVideoResp, err error) {
	return s.svr.UploadVideo(ctx, req)
}

func (s *Server) DownLoadVideo(ctx context.Context, req *videoApi.DownloadReq) (resp *videoApi.DownLoadResp, err error) {
	return s.svr.DownLoadVideo(ctx, req)
}

func (s *Server) GetVideoInfoById(ctx context.Context, req *videoApi.GetVideoInfoByIdReq) (resp *videoApi.GetVideoInfoByIdResp, err error) {
	return s.svr.GetVideoInfoById(ctx, req)
}

func (s *Server) CallbackOne(ctx context.Context, req *videoApi.RebackOneReq) (resp *videoApi.RebackOneResp, err error) {
	return s.svr.CallbackOne(ctx, req)
}

func (s *Server) CallbackTwo(ctx context.Context, req *videoApi.RebackTwoReq) (resp *videoApi.RebackTwoResp, err error) {
	return s.svr.CallbackTwo(ctx, req)
}

func (s *Server) PublishList(ctx context.Context, req *videoApi.PublishListReq) (resp *videoApi.PublishListResp, err error) {
	return s.svr.PublishList(ctx, req)
}

func (s *Server) SearchVideoByVideoTag(ctx context.Context, req *videoApi.SearchVideoByVideoTagReq) (resp *videoApi.SearchVideoByVideoTagResp, err error) {
	return s.svr.SearchVideoByVideoTag(ctx, req)
}

func (s *Server) InrcCommentCount(ctx context.Context, req *videoApi.InrcCommentCountReq) (resp *videoApi.InrcCommentCountResp, err error) {
	return s.svr.InrcCommentCount(ctx, req)
}
