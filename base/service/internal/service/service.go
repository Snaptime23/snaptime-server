package service

import (
	"context"
	"database/sql"
	"encoding/base64"
	"github.com/Snaptime23/snaptime-server/v2/base/internal/baseApi"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao/model"
	"github.com/Snaptime23/snaptime-server/v2/tools/errno"
	"github.com/Snaptime23/snaptime-server/v2/tools/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
	"time"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) UserRegister(ctx context.Context, req *baseApi.UserRegisterReq) (resp *baseApi.UserRegisterResp, err error) {
	resp = new(baseApi.UserRegisterResp)
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

func (s *Service) UserLogin(ctx context.Context, req *baseApi.UserLoginReq) (resp *baseApi.UserLoginResp, err error) {
	resp = new(baseApi.UserLoginResp)
	user, err := dao.GetUserByName(ctx, req.UserName)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return
	}
	resp.UserId = user.UserID
	resp.Token, err = jwt.GenToken(user.UserID, req.UserName)
	return
}

func (s *Service) UserInfo(ctx context.Context, req *baseApi.UserInfoReq) (resp *baseApi.UserInfoResp, err error) {
	user, err := dao.GetUserByName(ctx, req.UserId)
	if err != nil {
		return
	}
	resp = &baseApi.UserInfoResp{
		User: &baseApi.UserInfo{
			UserId:        user.UserID,
			UserName:      user.UserName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      0,
			Avatar:        user.Avatar,
			PublishNum:    user.VideoNum,
			FavouriteNum:  user.FavouriteNum,
		},
	}
	return
}

func (s *Service) PublishList(ctx context.Context, req *baseApi.PublishListReq) (resp *baseApi.PublishListResp, err error) {
	resp = new(baseApi.PublishListResp)
	return
}

func (s *Service) CreateComment(ctx context.Context, req *baseApi.CreateCommentReq) (resp *baseApi.CreateCommentResp, err error) {
	resp = new(baseApi.CreateCommentResp)
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
			RootID:      req.RootId,
			ParentID:    req.ParentId,
			PublishDate: time.Now().Format(time.DateTime),
		})
	} else { // delete
		err = dao.DeleteCommentByID(ctx, req.CommentId)
	}
	return
}

func generateToken(page int64) string {
	l := strconv.Itoa(int(page))
	s := []byte(l)
	token := base64.StdEncoding.EncodeToString(s)
	return token
}

func parseToken(token string) (left int64) {
	b, _ := base64.StdEncoding.DecodeString(token)
	s := strings.Split(string(b), "-")
	left, _ = strconv.ParseInt(s[0], 10, 64)
	return
}

func (s *Service) CommentList(ctx context.Context, req *baseApi.CommentListReq) (resp *baseApi.CommentListResp, err error) {
	resp = new(baseApi.CommentListResp)

	var page int64 = 1
	if req.Token != "" {
		page = parseToken(req.Token)
	}
	sum := dao.GetCommentCount(req.VideoId)
	if page*10 < sum {
		resp.HasNext = 1
		resp.NextPageToken = generateToken(page + 1)
	}

	ret, err := dao.GetPageQue(ctx, req.VideoId, page)
	list := make([]*baseApi.CommentInfo, 0)
	for _, val := range ret {
		user, err := dao.GetUserById(ctx, val.UserId)
		if err != nil {
			continue
		}
		chl := make([]*baseApi.CommentInfo, 0)
		for _, ch := range val.Children {
			chl = append(chl, &baseApi.CommentInfo{
				CommentId: ch.CommentID,
				User: &baseApi.UserInfo{
					UserId:   ch.UserId,
					UserName: ch.UserName,
					Avatar:   ch.UserAvatar,
				},
				VideoId:     ch.VideoId,
				Content:     ch.Content,
				PublishDate: ch.PublishDate,
				Replies:     nil,
			})
		}
		list = append(list, &baseApi.CommentInfo{
			CommentId: val.CommentID,
			User: &baseApi.UserInfo{
				UserId:        user.UserID,
				UserName:      user.UserName,
				FollowCount:   user.FollowerCount,
				FollowerCount: user.FollowCount,
				IsFollow:      0,
				Avatar:        user.Avatar,
				PublishNum:    user.VideoNum,
				FavouriteNum:  user.FavouriteNum,
			},
			VideoId:     val.VideoId,
			Content:     val.Content,
			PublishDate: val.PublishDate,
			Replies:     chl,
		})
	}
	resp.List = list
	return
}
