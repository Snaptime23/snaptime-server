package server

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/base/internal/baseApi"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/service"
)

type Server struct {
	svr *service.Service
	baseApi.UnimplementedBaseServiceServer
}

func NewServer() *Server {
	return &Server{
		svr:                            service.NewService(),
		UnimplementedBaseServiceServer: baseApi.UnimplementedBaseServiceServer{},
	}
}

func (s *Server) UserRegister(ctx context.Context, req *baseApi.UserRegisterReq) (resp *baseApi.UserRegisterResp, err error) {
	return s.svr.UserRegister(ctx, req)
}

func (s *Server) UserLogin(ctx context.Context, req *baseApi.UserLoginReq) (resp *baseApi.UserLoginResp, err error) {
	return s.svr.UserLogin(ctx, req)
}

func (s *Server) UserInfo(ctx context.Context, req *baseApi.UserInfoReq) (resp *baseApi.UserInfoResp, err error) {
	return s.svr.UserInfo(ctx, req)
}

func (s *Server) PublishList(ctx context.Context, req *baseApi.PublishListReq) (resp *baseApi.PublishListResp, err error) {
	return s.svr.PublishList(ctx, req)
}

func (s *Server) CreateComment(ctx context.Context, req *baseApi.CreateCommentReq) (resp *baseApi.CreateCommentResp, err error) {
	return s.svr.CreateComment(ctx, req)
}

func (s *Server) CommentList(ctx context.Context, req *baseApi.CommentListReq) (resp *baseApi.CommentListResp, err error) {
	return s.svr.CommentList(ctx, req)
}
