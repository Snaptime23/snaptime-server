package service

import (
	"context"
	"github.com/Snaptime23/snaptime-server/v2/base/internal/api"
	"google.golang.org/grpc"
)

type Service struct {
	baseClient api.BaseServiceClient
}

func NewService(conn *grpc.ClientConn) *Service {
	return &Service{baseClient: api.NewBaseServiceClient(conn)}
}

func (s *Service) UserRegister(username, password, confirmPassword string) (*api.UserRegisterResp, error) {
	return s.baseClient.UserRegister(context.Background(), &api.UserRegisterReq{
		UserName:        username,
		Password:        password,
		ConfirmPassword: confirmPassword,
	})
}

func (s *Service) UserLogin(UserName, Password string) (*api.UserLoginResp, error) {
	return s.baseClient.UserLogin(context.Background(), &api.UserLoginReq{
		UserName: UserName,
		Password: Password,
	})
}

func (s *Service) UserInfo(UserId string) (*api.UserInfoResp, error) {
	return s.baseClient.UserInfo(context.Background(), &api.UserInfoReq{
		UserId: UserId,
	})
}

func (s *Service) PublishList(UserId string) (*api.PublishListResp, error) {
	return s.baseClient.PublishList(context.Background(), &api.PublishListReq{
		UserId: UserId,
	})
}

func (s *Service) CreateComment(ctx context.Context, req *api.CreateCommentReq) (*api.CreateCommentResp, error) {
	return s.baseClient.CreateComment(ctx, req)
}

func (s *Service) CommentList(ctx context.Context, req *api.CommentListReq) (*api.CommentListResp, error) {
	return s.baseClient.CommentList(ctx, req)
}
