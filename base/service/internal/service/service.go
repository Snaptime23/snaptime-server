package service

import (
	"context"
	"database/sql"
	"github.com/Snaptime23/snaptime-server/v2/base/internal/api"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao/model"
	"github.com/Snaptime23/snaptime-server/v2/tools/errno"
	"github.com/Snaptime23/snaptime-server/v2/tools/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
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

func (s *Service) CreateComment(ctx context.Context, req *api.CreateCommentReq) (resp *api.CreateCommentResp, err error) {
	resp = new(api.CreateCommentResp)
	if req.ActionType == 0 { // create
		resp.CommentId = uuid.NewString()
		err = dao.CreateComment(ctx, &model.Comment{
			CommentID:   resp.CommentId,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
			DeletedAt:   sql.NullTime{},
			UserId:      req.UserID,
			VideoId:     req.VideoId,
			Content:     req.Content,
			PublishDate: time.Now().Format(time.DateTime),
		})
	} else { // delete
		err = dao.DeleteCommentByID(ctx, req.CommentId)
	}
	return
}

func (s *Service) CommentList(ctx context.Context, req *api.CommentListReq) (resp *api.CommentListResp, err error) {
	resp = new(api.CommentListResp)
	ret, err := dao.GetCommentList(ctx, req.VideoId)
	list := make([]*api.CommentInfo, 0)
	for _, val := range ret {
		user, err := dao.GetUserById(ctx, val.UserId)
		if err != nil {
			continue
		}
		list = append(list, &api.CommentInfo{
			CommentId: val.CommentID,
			User: &api.UserInfo{
				UserId:          user.ID,
				UserName:        user.UserName,
				FollowCount:     user.FollowerCount,
				FollowerCount:   user.FollowCount,
				IsFollow:        0,
				Avatar:          user.Avatar,
				PublishNum:      user.PublishNum,
				FavouriteNum:    user.FavouriteNum,
				LikeNum:         user.LikeNum,
				ReceivedLikeNum: user.ReceivedLikeNum,
			},
			VideoId:     val.VideoId,
			Content:     val.Content,
			PublishDate: val.PublishDate,
		})
	}
	resp.List = list
	return
}
