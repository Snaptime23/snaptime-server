package service

import (
	"context"
	"database/sql"
	"encoding/base64"
	"github.com/Snaptime23/snaptime-server/v2/base/rpc_pb/baseApi"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/dao/model"
	"github.com/Snaptime23/snaptime-server/v2/tools/errno"
	"github.com/Snaptime23/snaptime-server/v2/tools/jwt"
	"github.com/Snaptime23/snaptime-server/v2/video/rpc_pb/videoApi"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
	"strings"
	"time"
)

type Service struct {
	videoClient videoApi.VideoServiceClient
}

func NewService() *Service {
	conn, err := grpc.Dial(":9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &Service{
		videoClient: videoApi.NewVideoServiceClient(conn),
	}
}

func (s *Service) UserRegister(ctx context.Context, req *baseApi.UserRegisterReq) (resp *baseApi.UserRegisterResp, err error) {
	resp = new(baseApi.UserRegisterResp)
	user, err := dao.GetUserByName(ctx, req.UserName)
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
	user, err := dao.GetUserById(ctx, req.UserId)
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
			PublishDate: time.Now().Unix(),
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

	ret, err := dao.GetPageQue(ctx, req.VideoId, req.RootID, page)
	list := make([]*baseApi.CommentInfo, 0)
	for _, val := range ret {
		user, err := dao.GetUserById(ctx, val.UserId)
		if err != nil {
			continue
		}
		hasLike, err := dao.HasLikeComment(ctx, val.CommentID, user.UserID)
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
			Replies:     dao.GetChildrenNum(req.VideoId, val.CommentID),
			LikeCount:   val.LikeCount,
			HasLike:     hasLike,
		})
	}
	resp.List = list
	return
}

func (s *Service) LikeComment(ctx context.Context, req *baseApi.LikeCommentReq) (resp *baseApi.LikeCommentResp, err error) {
	resp = new(baseApi.LikeCommentResp)
	err = dao.UpdateCommentLikeCount(ctx, req.CommentID, req.UserId, req.ActionType)

	return
}

func (s *Service) LikeVideoAction(ctx context.Context, req *baseApi.LikeVideoActionReq) (resp *baseApi.LikeVideoActionResp, err error) {
	resp = new(baseApi.LikeVideoActionResp)
	err = dao.UpdateAndInsertLikeRecord(ctx, req.UserId, req.VideoId, req.ActionType)
	return
}

func (s *Service) VideoLikeList(ctx context.Context, req *baseApi.VideoLikeListReq) (resp *baseApi.VideoLikeListResp, err error) {
	resp = new(baseApi.VideoLikeListResp)
	resp.VideoList = make([]*baseApi.VideoInfo, 0)
	videoIDS, err := dao.GetUserLikeRecords(ctx, req.UserId)
	for _, videoID := range videoIDS {
		video, err := s.videoClient.GetVideoInfoById(ctx, &videoApi.GetVideoInfoByIdReq{
			VideoId: videoID,
		})
		if err != nil {
			continue
		}
		resp.VideoList = append(resp.VideoList, &baseApi.VideoInfo{
			VideoId: videoID,
			UserInfo: &baseApi.UserInfo{
				UserId:          video.Video.Author.UserId,
				UserName:        video.Video.Author.UserName,
				FollowCount:     video.Video.Author.FollowerCount,
				FollowerCount:   video.Video.Author.FollowerCount,
				IsFollow:        0,
				Avatar:          video.Video.Author.Avatar,
				PublishNum:      video.Video.Author.PublishNum,
				FavouriteNum:    video.Video.Author.FavouriteNum,
				LikeNum:         video.Video.Author.LikeNum,
				ReceivedLikeNum: video.Video.Author.ReceivedLikeNum,
			},
			FavoriteCount: video.Video.FavoriteCount,
			CommentCount:  video.Video.CommentCount,
			IsFavorite:    0,
			Title:         video.Video.Title,
		})
	}
	return
}

func (s *Service) FollowList(ctx context.Context, req *baseApi.FollowListReq) (resp *baseApi.FollowListResp, err error) {
	resp = new(baseApi.FollowListResp)
	resp.User = make([]*baseApi.UserInfo, 0)
	folloList, err := dao.GetFollowList(ctx, req.UserId)
	if err != nil {
		return
	}
	for _, val := range folloList {
		user, err := dao.GetUserById(ctx, val)
		if err != nil {
			continue
		}
		resp.User = append(resp.User, &baseApi.UserInfo{
			UserId:          user.UserID,
			UserName:        user.UserName,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        0,
			Avatar:          user.Avatar,
			PublishNum:      user.VideoNum,
			FavouriteNum:    user.FavouriteNum,
			LikeNum:         user.FavouriteNum,
			ReceivedLikeNum: 0,
		})
	}
	return
}

func (s *Service) FollowerList(ctx context.Context, req *baseApi.FollowerListReq) (resp *baseApi.FollowerListResp, err error) {
	resp = new(baseApi.FollowerListResp)
	resp.User = make([]*baseApi.UserInfo, 0)
	folloList, err := dao.GetFollowerList(ctx, req.UserId)
	if err != nil {
		return
	}
	for _, val := range folloList {
		user, err := dao.GetUserById(ctx, val)
		if err != nil {
			continue
		}
		resp.User = append(resp.User, &baseApi.UserInfo{
			UserId:          user.UserID,
			UserName:        user.UserName,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        0,
			Avatar:          user.Avatar,
			PublishNum:      user.VideoNum,
			FavouriteNum:    user.FavouriteNum,
			LikeNum:         user.FavouriteNum,
			ReceivedLikeNum: 0,
		})
	}
	return
}
