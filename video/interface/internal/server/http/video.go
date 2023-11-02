package http

import (
	"github.com/Snaptime23/snaptime-server/v2/video/interface/internal/service"
	"google.golang.org/grpc"
)

type HttpServer struct {
	svr *service.Service
}

func NewServer(conn *grpc.ClientConn) *HttpServer {
	return &HttpServer{svr: service.NewService(conn)}
}
