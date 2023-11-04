package http

import (
	"context"
	"fmt"
	"github.com/Snaptime23/snaptime-server/v2/base/interface/internal/service"
	"github.com/Snaptime23/snaptime-server/v2/base/rpc_pb/baseApi"
	"github.com/Snaptime23/snaptime-server/v2/tools"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type HttpServer struct {
	svr *service.Service
}

func NewServer(conn *grpc.ClientConn) *HttpServer {
	return &HttpServer{svr: service.NewService(conn)}
}

func (s *HttpServer) UserRegister(c *gin.Context) {
	arg := new(struct {
		UserName        string `json:"user_name"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirm_password"`
	})
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	fmt.Println(arg)
	resp, err := s.svr.UserRegister(arg.UserName, arg.Password, arg.ConfirmPassword)
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) UserLogin(c *gin.Context) {
	arg := new(struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	})
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	resp, err := s.svr.UserLogin(arg.UserName, arg.Password)
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) UserInfo(c *gin.Context) {
	arg := new(struct {
		UserId string `json:"user_id"`
	})
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	resp, err := s.svr.UserInfo(arg.UserId)
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) PublishList(c *gin.Context) {
	arg := new(struct {
		UserId string `json:"user_id"`
	})
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	resp, err := s.svr.PublishList(arg.UserId)
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) CreateComment(c *gin.Context) {
	arg := new(struct {
		VideoId    string `json:"video_id"`
		ActionType int64  `json:"action_type"`
		Content    string `json:"content"`
		CommentId  string `json:"comment_id"`
		RootId     string `json:"root_id"`
		ParentId   string `json:"parent_id"`
	})
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	//fmt.Println("userID = ", c.Get("user_id"))
	userID, _ := c.Get("user_id")
	resp, err := s.svr.CreateComment(context.Background(), &baseApi.CreateCommentReq{
		VideoId:    arg.VideoId,
		ActionType: arg.ActionType,
		Content:    arg.Content,
		CommentId:  arg.CommentId,
		UserID:     userID.(string),
		RootId:     arg.RootId,
		ParentId:   arg.ParentId,
	})
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) CommentList(c *gin.Context) {
	arg := new(struct {
		VideoId  string
		RootId   string
		ParentID string
		Token    string
	})
	arg.VideoId = c.Query("video_id")
	arg.Token = c.Query("token")
	arg.RootId = c.Query("root_id")
	arg.ParentID = c.Query("parent_id")
	resp, err := s.svr.CommentList(context.Background(), &baseApi.CommentListReq{
		VideoId:  arg.VideoId,
		Token:    arg.Token,
		ParentID: arg.ParentID,
		RootID:   arg.RootId,
	})
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) LikeVideoAction(c *gin.Context) {
	arg := new(struct {
		UserId     string `json:"user_id"`
		VideoId    string `json:"video_id"`
		ActionType int64  `json:"action_type"`
	})
	userID, _ := c.Get("user_id")
	arg.UserId = userID.(string)
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	resp, err := s.svr.LikeVideoAction(context.Background(), &baseApi.LikeVideoActionReq{
		UserId:     arg.UserId,
		VideoId:    arg.VideoId,
		ActionType: arg.ActionType,
	})
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) VideoLikeList(c *gin.Context) {
	arg := new(struct {
		UserId string `json:"user_id"`
	})
	userID, _ := c.Get("user_id")
	arg.UserId = userID.(string)
	resp, err := s.svr.VideoLikeList(context.Background(), &baseApi.VideoLikeListReq{
		UserId: arg.UserId,
	})
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) LikeComment(c *gin.Context) {
	arg := new(struct {
		CommentID  string `json:"comment_id"`
		ActionType int64  `json:"action_type"`
	})
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	resp, err := s.svr.LikeComment(context.Background(), &baseApi.LikeCommentReq{
		CommentID:  arg.CommentID,
		ActionType: arg.ActionType,
	})
	tools.HandleErrOrResp(c, resp, err)
}
