package service

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/base/internal/api"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao"
	"github.com/Snaptime23/snaptime-server/v2/tools/errno"
	"github.com/Snaptime23/snaptime-server/v2/tools/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) UserRegister(ctx context.Context, req *api.UserRegisterReq) (resp *api.UserRegisterResp, err error) {
	resp = new(api.UserRegisterResp)
	user, err := dao.GetUserByName(ctx, req.UserName)
	if req.Password != req.ConfirmPassword {
		return resp, errno.NewErrNo("两次密码不一致")
	}
	if user != nil {
		return resp, errno.NewErrNo("用户已存在")
	}
	EncodePassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	resp.UserId, err = dao.CreateUser(ctx, req.UserName, string(EncodePassword))
	if err != nil {
		return
	}
	resp.Token, err = jwt.GenToken(resp.UserId, req.UserName)
	return
}

func (s *Service) UserLogin(ctx context.Context, req *api.UserLoginReq) (resp *api.UserLoginResp, err error) {
	resp = new(api.UserLoginResp)
	user, err := dao.GetUserByName(ctx, req.UserName)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return
	}
	resp.UserId = user.ID
	resp.Token, err = jwt.GenToken(user.ID, req.UserName)
	return
}

func (s *Service) UserInfo(ctx context.Context, req *api.UserInfoReq) (resp *api.UserInfoResp, err error) {
	user, err := dao.GetUserByName(ctx, req.UserId)
	if err != nil {
		return
	}
	resp = &api.UserInfoResp{
		User: &api.UserInfo{
			UserId:          user.ID,
			UserName:        user.UserName,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        0,
			Avatar:          user.Avatar,
			PublishNum:      user.PublishNum,
			FavouriteNum:    user.FavouriteNum,
			LikeNum:         user.LikeNum,
			ReceivedLikeNum: user.ReceivedLikeNum,
		},
	}
	return
}

func (s *Service) PublishList(ctx context.Context, req *api.PublishListReq) (resp *api.PublishListResp, err error) {
	resp = new(api.PublishListResp)
	return
}
