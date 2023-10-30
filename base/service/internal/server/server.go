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

func NewServer() *Server {
	return &Server{
		svr:                            service.NewService(),
		UnimplementedBaseServiceServer: api.UnimplementedBaseServiceServer{},
	}
}

func (s *Server) UserRegister(ctx context.Context, req *api.UserRegisterReq) (resp *api.UserRegisterResp, err error) {
	return s.svr.UserRegister(ctx, req)
}

func (s *Server) UserLogin(ctx context.Context, req *api.UserLoginReq) (resp *api.UserLoginResp, err error) {
	return s.svr.UserLogin(ctx, req)
}

func (s *Server) UserInfo(ctx context.Context, req *api.UserInfoReq) (resp *api.UserInfoResp, err error) {
	return s.svr.UserInfo(ctx, req)
}

func (s *Server) PublishList(ctx context.Context, req *api.PublishListReq) (resp *api.PublishListResp, err error) {
	return s.svr.PublishList(ctx, req)
}

func (s *Server) CreateComment(ctx context.Context, req *api.CreateCommentReq) (resp *api.CreateCommentResp, err error) {
	return s.svr.CreateComment(ctx, req)
}

func (s *Server) CommentList(ctx context.Context, req *api.CommentListReq) (resp *api.CommentListResp, err error) {
	return s.svr.CommentList(ctx, req)
}
