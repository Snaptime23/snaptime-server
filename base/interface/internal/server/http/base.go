package http

import (
	"github.com/Snaptime23/snaptime-server/v2/base/interface/internal/service"
	"github.com/Snaptime23/snaptime-server/v2/tools"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type HttpServer struct {
	svr *service.Service
}

func NewServer(conn *grpc.ClientConn) *HttpServer {
	return &HttpServer{svr: service.NewService(conn)}
}

func (s *HttpServer) UserRegister(c *gin.Context) {
	arg := new(struct {
		UserName        string
		Password        string
		ConfirmPassword string
	})
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	resp, err := s.svr.UserRegister(arg.UserName, arg.Password, arg.ConfirmPassword)
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) UserLogin(c *gin.Context) {
	arg := new(struct {
		UserName string
		Password string
	})
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	resp, err := s.svr.UserLogin(arg.UserName, arg.Password)
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) UserInfo(c *gin.Context) {
	arg := new(struct {
		UserId string
	})
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	resp, err := s.svr.UserInfo(arg.UserId)
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) PublishList(c *gin.Context) {
	arg := new(struct {
		UserId string
	})
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	resp, err := s.svr.PublishList(arg.UserId)
	tools.HandleErrOrResp(c, resp, err)
}
