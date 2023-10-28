package service

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/base/internal/api"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) UserRegister(ctx context.Context, req *api.UserRegisterReq) (resp *api.UserRegisterResp, err error) {
	resp = new(api.UserRegisterResp)
	return
}

func (s *Service) UserLogin(ctx context.Context, req *api.UserLoginReq) (resp *api.UserLoginResp, err error) {
	resp = new(api.UserLoginResp)
	return
}

func (s *Service) UserInfo(ctx context.Context, req *api.UserInfoReq) (resp *api.UserInfoResp, err error) {
	resp = new(api.UserInfoResp)
	return
}

func (s *Service) PublishList(ctx context.Context, req *api.PublishListReq) (resp *api.PublishListResp, err error) {
	resp = new(api.PublishListResp)
	return
}
