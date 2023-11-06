package service

import (
	"context"
	"database/sql"
	"encoding/base64"
	"github.com/Snaptime23/snaptime-server/v2/base/rpc_pb/baseApi"
	"github.com/Snaptime23/snaptime-server/v2/base/service/internal/cache"
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
		_, err = s.videoClient.InrcCommentCount(ctx, &videoApi.InrcCommentCountReq{
			VideoId: req.VideoId,
			Count:   1,
		})
	} else { // delete
		err = dao.DeleteCommentByID(ctx, req.CommentId)
		_, err = s.videoClient.InrcCommentCount(ctx, &videoApi.InrcCommentCountReq{
			VideoId: req.VideoId,
			Count:   -1,
		})
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
	if cache.LikeIsExists(ctx, req.UserId) == 0 {
		likeList, err := dao.GetUserLikeRecords(ctx, req.UserId)
		if err != nil {
			return resp, err
		}
		if len(likeList) > 0 {
			kv := make([]string, 0)
			for _, videoId := range likeList {
				kv = append(kv, videoId)
				kv = append(kv, "1")
			}
			if !cache.SetFavoriteList(ctx, req.UserId, kv...) {
				return resp, errno.NewErrNo("Redis设置用户点赞视频缓存出错")
			}
		}
	}
	if cache.UserIsExists(ctx, req.UserId) == 0 {
		user, err := dao.GetUserById(ctx, req.UserId)
		if err != nil {
			return resp, err
		}
		if !cache.SetUserInfo(ctx, user) {
			return resp, errno.NewErrNo("Redis缓存用户信息出错")
		}
	}
	if cache.VideoIsExists(ctx, req.VideoId) == 0 {
		video, err := s.videoClient.GetVideoInfoById(ctx, &videoApi.GetVideoInfoByIdReq{
			VideoId: req.VideoId,
		})
		if err != nil {
			return nil, err
		}
		if !cache.SetVideoMessage(ctx, &videoApi.VideoInfo{
			VideoID:       video.Video.VideoID,
			Author:        video.Video.Author,
			PlayUrl:       video.Video.PlayUrl,
			CoverUrl:      video.Video.CoverUrl,
			FavoriteCount: video.Video.FavoriteCount,
			CommentCount:  video.Video.CommentCount,
			Title:         video.Video.Title,
			IsEncoding:    video.Video.IsEncoding,
		}) {
			return nil, errno.NewErrNo("Redis缓存视频信息出错")
		}
	}
	res := cache.GetVideoFields(ctx, req.VideoId, "user_id")
	authorID, _ := res[0].(string)
	if cache.UserIsExists(ctx, authorID) == 0 {
		user, err := dao.GetUserById(ctx, authorID)
		if err != nil {
			return resp, err
		}
		if !cache.SetUserInfo(ctx, user) {
			return resp, errno.NewErrNo("Redis缓存用户信息出错")
		}
	}

	var action int64
	like := cache.IsLike(ctx, req.UserId, req.VideoId)
	if like && req.GetActionType() == 1 {
		action = -1
	} else if !like && req.GetActionType() == 0 {
		action = 1
	} else {
		return
	}

	// 更新点赞列表
	if !cache.FavoriteAction(ctx, req.UserId, req.VideoId, action) {
		return resp, errno.NewErrNo("更新点赞列表失败")
	}
	// 更新用户点赞数
	if !cache.IncrUserField(ctx, req.UserId, "favorite_count", action) {
		return resp, errno.NewErrNo("更新用户点赞数")
	}
	// 更新作者获赞数
	if !cache.IncrUserField(ctx, authorID, "total_favorited", action) {
		return resp, errno.NewErrNo("更新点赞列表失败")
	}
	// 更新视频获赞数
	if !cache.IncrVideoField(ctx, req.VideoId, "favorite_count", action) {
		return resp, errno.NewErrNo("更新视频获赞数")
	}

	return
}

func (s *Service) VideoLikeList(ctx context.Context, req *baseApi.VideoLikeListReq) (resp *baseApi.VideoLikeListResp, err error) {
	resp = new(baseApi.VideoLikeListResp)
	var likeList []string
	if cache.LikeIsExists(ctx, req.UserId) == 0 {
		likeList, err = dao.GetUserLikeRecords(ctx, req.UserId)
		if err != nil {
			return resp, err
		}
		if len(likeList) > 0 {
			kv := make([]string, 0)
			for _, videoId := range likeList {
				kv = append(kv, videoId)
				kv = append(kv, "1")
			}
			if !cache.SetFavoriteList(ctx, req.UserId, kv...) {
				return resp, errno.NewErrNo("Redis设置用户点赞视频缓存出错")
			}
		}
	} else {
		likeList = cache.GetAllUserLikes(ctx, req.UserId)
	}

	var videoList []*videoApi.VideoInfo
	for _, vid := range likeList {
		var video *videoApi.GetVideoInfoByIdResp
		if cache.VideoIsExists(ctx, vid) == 0 {
			video, err = s.videoClient.GetVideoInfoById(ctx, &videoApi.GetVideoInfoByIdReq{
				VideoId: vid,
			})
			if err != nil {
				return nil, err
			}
			cache.SetVideoMessage(ctx, video.Video)
		} else {
			video.Video, err = cache.GetVideoMessage(ctx, vid)
		}

		var user *model.User
		if cache.UserIsExists(ctx, video.Video.Author.UserId) == 0 {
			user, err = dao.GetUserById(ctx, video.Video.Author.UserId)
			if err != nil {
				return nil, err
			}
			cache.SetUserInfo(ctx, user)
		} else {
			user, err = cache.GetUserInfo(ctx, video.Video.Author.UserId)
		}

		retVideo := &videoApi.VideoInfo{
			VideoID: vid,
			Author: &videoApi.UserInfo{
				UserId:        video.Video.Author.UserId,
				UserName:      user.UserName,
				FollowerCount: 0,
				IsFollow:      0,
			},
			PlayUrl:       video.Video.PlayUrl,
			CoverUrl:      video.Video.CoverUrl,
			FavoriteCount: video.Video.FavoriteCount,
			CommentCount:  0,
			IsFavorite:    0,
			Title:         video.Video.Title,
		}
		videoList = append(videoList, retVideo)
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

func (s *Service) Follow(ctx context.Context, req *baseApi.FollowReq) (resp *baseApi.FollowResp, err error) {
	resp = new(baseApi.FollowResp)
	err = dao.Follow(ctx, req.UserId, req.ToUserId, req.ActionType)
	return
}
