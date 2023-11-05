package service

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/base/rpc_pb/baseApi"
	"github.com/Snaptime23/snaptime-server/v2/video/rpc_pb/videoApi"
	"google.golang.org/grpc"
)

type Service struct {
	baseClient  baseApi.BaseServiceClient
	videoClient videoApi.VideoServiceClient
}

func NewService(connBase, connVideo *grpc.ClientConn) *Service {
	return &Service{
		baseClient:  baseApi.NewBaseServiceClient(connBase),
		videoClient: videoApi.NewVideoServiceClient(connVideo),
	}
}

func (s *Service) UserRegister(username, password string) (*baseApi.UserRegisterResp, error) {
	return s.baseClient.UserRegister(context.Background(), &baseApi.UserRegisterReq{
		UserName: username,
		Password: password,
	})
}

func (s *Service) UserLogin(UserName, Password string) (*baseApi.UserLoginResp, error) {
	return s.baseClient.UserLogin(context.Background(), &baseApi.UserLoginReq{
		UserName: UserName,
		Password: Password,
	})
}

func (s *Service) UserInfo(UserId string) (*baseApi.UserInfoResp, error) {
	return s.baseClient.UserInfo(context.Background(), &baseApi.UserInfoReq{
		UserId: UserId,
	})
}

func (s *Service) PublishList(UserId string) (*videoApi.PublishListResp, error) {
	return s.videoClient.PublishList(context.Background(), &videoApi.PublishListReq{
		UserId: UserId,
	})
}

func (s *Service) CreateComment(ctx context.Context, req *baseApi.CreateCommentReq) (*baseApi.CreateCommentResp, error) {
	return s.baseClient.CreateComment(ctx, req)
}

func (s *Service) CommentList(ctx context.Context, req *baseApi.CommentListReq) (*baseApi.CommentListResp, error) {
	return s.baseClient.CommentList(ctx, req)
}

func (s *Service) LikeVideoAction(ctx context.Context, req *baseApi.LikeVideoActionReq) (*baseApi.LikeVideoActionResp, error) {
	return s.baseClient.LikeVideoAction(ctx, req)
}

func (s *Service) VideoLikeList(ctx context.Context, req *baseApi.VideoLikeListReq) (*baseApi.VideoLikeListResp, error) {
	return s.baseClient.VideoLikeList(ctx, req)
}

func (s *Service) LikeComment(ctx context.Context, req *baseApi.LikeCommentReq) (resp *baseApi.LikeCommentResp, err error) {
	return s.baseClient.LikeComment(ctx, req)
}

func (s *Service) UploadVideo(ctx context.Context, req *videoApi.UploadVideoReq) (resp *videoApi.UploadVideoResp, err error) {
	return s.videoClient.UploadVideo(ctx, req)
}

func (s *Service) DownLoadVideo(ctx context.Context, req *videoApi.DownloadReq) (resp *videoApi.DownLoadResp, err error) {
	return s.videoClient.DownLoadVideo(ctx, req)
}

func (s *Service) CallbackOne(ctx context.Context, req *videoApi.RebackOneReq) (resp *videoApi.RebackOneResp, err error) {
	return s.videoClient.CallbackOne(ctx, req)
}

func (s *Service) CallbackTwo(ctx context.Context, req *videoApi.RebackTwoReq) (resp *videoApi.RebackTwoResp, err error) {
	return s.videoClient.CallbackTwo(ctx, req)
}

func (s *Service) FollowList(ctx context.Context, req *baseApi.FollowListReq) (resp *baseApi.FollowListResp, err error) {
	return s.baseClient.FollowList(ctx, req)
}

func (s *Service) FollowerList(ctx context.Context, req *baseApi.FollowerListReq) (resp *baseApi.FollowerListResp, err error) {
	return s.baseClient.FollowerList(ctx, req)
}

func (s *Service) Follow(ctx context.Context, req *baseApi.FollowReq) (resp *baseApi.FollowResp, err error) {
	return s.baseClient.Follow(ctx, req)
}
