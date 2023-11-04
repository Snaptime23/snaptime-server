package http

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/tools"
	"github.com/Snaptime23/snaptime-server/v2/video/interface/internal/service"
	"github.com/Snaptime23/snaptime-server/v2/video/internal/videoApi"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type HttpServer struct {
	svr       *service.Service
	videoChan chan string
}

func NewServer(conn *grpc.ClientConn) *HttpServer {
	return &HttpServer{
		svr:       service.NewService(conn),
		videoChan: make(chan string, 20),
	}
}

func (s *HttpServer) UpLoadVideo(c *gin.Context) {
	resp, err := s.svr.UploadVideoToken(context.Background(), &videoApi.UploadVideoReq{})
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) DownLoadVideo(c *gin.Context) {
	resp, err := s.svr.DownLoadVideoToken(context.Background(), &videoApi.DownloadReq{})
	tools.HandleErrOrResp(c, resp, err)
}
