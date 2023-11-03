package server

import (
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