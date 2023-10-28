package server

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/base/internal/api"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/service"
)

type Server struct {
	svr *service.Service
	api.UnimplementedBaseServiceServer
}

func (s *Server) UserRegister(ctx context.Context, req *api.UserRegisterReq) (resp *api.UserRegisterResp, err error) {
	return
}

func (s *Server) UserLogin(ctx context.Context, req *api.UserLoginReq) (resp *api.UserLoginResp, err error) {
	return
}

func (s *Server) UserInfo(ctx context.Context, req *api.UserInfoReq) (resp *api.UserInfoResp, err error) {
	return
}

func (s *Server) PublishList(ctx context.Context, req *api.PublishListReq) (resp *api.PublishListResp, err error) {
	return
}
