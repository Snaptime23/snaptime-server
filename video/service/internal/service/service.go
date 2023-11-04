package service

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/base/rpc_pb/baseApi"
	"github.com/Snaptime23/snaptime-server/v2/video/rpc_pb/videoApi"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/dao"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/service/downloadToken"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/service/uploadToken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	baseClient baseApi.BaseServiceClient
}

func NewService() *Service {
	conn, err := grpc.Dial(":9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &Service{
		baseClient: baseApi.NewBaseServiceClient(conn),
	}
}

func (s *Service) VideoFeed(ctx context.Context, req *videoApi.VideoFeedReq) (resp *videoApi.VideoFeedResp, err error) {
	return
}

func (s *Service) UploadVideoToken(ctx context.Context, req *videoApi.UploadVideoReq) (resp *videoApi.UploadVideoResp, err error) {
	resp = new(videoApi.UploadVideoResp)
	resp.Token = downloadToken.GetToken()
	return
}

func (s *Service) DownLoadVideoToken(ctx context.Context, req *videoApi.DownloadReq) (resp *videoApi.DownLoadResp, err error) {
	resp = new(videoApi.DownLoadResp)
	resp.Token = uploadToken.GetToken()
	return
}

func (s *Service) GetVideoInfoById(ctx context.Context, req *videoApi.GetVideoInfoByIdReq) (resp *videoApi.GetVideoInfoByIdResp, err error) {
	resp = new(videoApi.GetVideoInfoByIdResp)
	video, err := dao.GetVideoByVideoId(ctx, req.VideoId)
	if err != nil {
		return
	}
	user, err := s.baseClient.UserInfo(ctx, &baseApi.UserInfoReq{
		UserId: video.CreateUserId,
	})
	resp.Video = &videoApi.VideoInfo{
		VideoID: video.VideoID,
		Author: &videoApi.UserInfo{
			UserId:          user.User.UserId,
			UserName:        user.User.UserName,
			FollowCount:     user.User.FollowCount,
			FollowerCount:   user.User.FollowerCount,
			IsFollow:        0,
			Avatar:          user.User.Avatar,
			PublishNum:      user.User.PublishNum,
			FavouriteNum:    user.User.FavouriteNum,
			LikeNum:         user.User.LikeNum,
			ReceivedLikeNum: user.User.ReceivedLikeNum,
		},
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavouriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    0,
		Title:         video.VideoName,
	}
	return
}
