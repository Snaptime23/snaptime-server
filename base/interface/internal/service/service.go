package service

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/base/rpc_pb/baseApi"
	"google.golang.org/grpc"
)

type Service struct {
	baseClient baseApi.BaseServiceClient
}

func NewService(conn *grpc.ClientConn) *Service {
	return &Service{baseClient: baseApi.NewBaseServiceClient(conn)}
}

func (s *Service) UserRegister(username, password, confirmPassword string) (*baseApi.UserRegisterResp, error) {
	return s.baseClient.UserRegister(context.Background(), &baseApi.UserRegisterReq{
		UserName:        username,
		Password:        password,
		ConfirmPassword: confirmPassword,
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

func (s *Service) PublishList(UserId string) (*baseApi.PublishListResp, error) {
	return s.baseClient.PublishList(context.Background(), &baseApi.PublishListReq{
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
