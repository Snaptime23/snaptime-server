package http

import (
	"context"
	"fmt"
	"github.com/Snaptime23/snaptime-server/v2/base/interface/internal/service"
	"github.com/Snaptime23/snaptime-server/v2/base/rpc_pb/baseApi"
	"github.com/Snaptime23/snaptime-server/v2/tools"
	"github.com/Snaptime23/snaptime-server/v2/video/rpc_pb/videoApi"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"strings"
)

type HttpServer struct {
	svr *service.Service
}

func NewServer(connBase, connVideo *grpc.ClientConn) *HttpServer {
	return &HttpServer{
		svr: service.NewService(connBase, connVideo),
	}
}

func (s *HttpServer) UserRegister(c *gin.Context) {
	arg := new(struct {
		UserName    string `json:"user_name"`
		Password    string `json:"password"`
		Avatar      string `json:"avatar"`
		Description string `json:"description"`
		Email       string `json:"email"`
	})
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	fmt.Println(arg)
	resp, err := s.svr.UserRegister(arg.UserName, arg.Password, arg.Avatar, arg.Description, arg.Email)
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
	arg.UserId = c.Query("user_id")
	resp, err := s.svr.UserInfo(arg.UserId)
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) PublishList(c *gin.Context) {
	arg := new(struct {
		UserId string `json:"user_id"`
	})
	arg.UserId = c.Query("user_id")
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
	if arg.VideoId == "" {
		tools.HandleErrOrResp(c, nil, nil)
	}
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
	userId, _ := c.Get("user_id")
	resp, err := s.svr.LikeComment(context.Background(), &baseApi.LikeCommentReq{
		CommentID:  arg.CommentID,
		ActionType: arg.ActionType,
		UserId:     userId.(string),
	})
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) UpLoadVideo(c *gin.Context) {
	arg := new(struct {
		Title         string   `json:"title"`
		Description   string   `json:"description"`
		VideoTags     []string `json:"video_tags"`
		FileExtension string   `json:"file_extension"`
		VideoId       string   `json:"video_id"`
	})
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	fmt.Println(arg)
	// return user_upload/{user_uuid}/{video_uuid.file_extension}
	userId, _ := c.Get("user_id")
	resp, err := s.svr.UploadVideo(context.Background(), &videoApi.UploadVideoReq{
		Description:   arg.Description,
		Title:         arg.Title,
		VideoTag:      arg.VideoTags,
		FileExtension: arg.FileExtension,
		UserId:        userId.(string),
		VideoId:       arg.VideoId,
	})
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) DownLoadVideo(c *gin.Context) {
	resp, err := s.svr.DownLoadVideo(context.Background(), &videoApi.DownloadReq{})
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) CallbackOne(c *gin.Context) {
	arg := new(struct {
		Input struct {
			KodoFile struct {
				Bucket string `json:"bucket"`
				Key    string `json:"key"`
			} `json:"kodo_file"`
		} `json:"input"`
		Ops []struct {
			ID  string `json:"id"`
			Fop struct {
				Result struct {
					KodoFile struct {
						Bucket string `json:"bucket"`
						Key    string `json:"key"`
					} `json:"kodo_file"`
				} `json:"result"`
			} `json:"fop,omitempty"`
		}
	})
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	tmp := strings.Split(arg.Input.KodoFile.Key, ".")
	coverUrl := ""
	for _, val := range arg.Ops {
		if val.ID == "node4_saveas" {
			coverUrl = val.Fop.Result.KodoFile.Key
			break
		}
	}
	if len(tmp) > 0 {
		key := tmp[0]
		resp, err := s.svr.CallbackOne(context.Background(), &videoApi.RebackOneReq{
			Title:    key,
			CoverUrl: coverUrl,
		})
		tools.HandleErrOrResp(c, resp, err)
	}
}

func (s *HttpServer) CallbackTwo(c *gin.Context) {
	arg := new(struct {
		Input struct {
			KodoFile struct {
				Bucket string `json:"bucket"`
				Key    string `json:"key"`
			} `json:"kodo_file"`
		} `json:"input"`

		Ops []struct {
			ID  string `json:"id"`
			Fop struct {
				Result struct {
					KodoFile struct {
						Bucket string `json:"bucket"`
						Key    string `json:"key"`
						Hash   string `json:"hash"`
					} `json:"kodo_file"`
				} `json:"result"`
			} `json:"fop,omitempty"`
		}
	})
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	tmp := strings.Split(arg.Input.KodoFile.Key, ".")
	ResourceKey := make([]string, 0)
	VideoId := ""
	if len(tmp) > 0 {
		tmp = strings.Split(tmp[0], "/")
		VideoId = tmp[2]
	}
	Quality := make([]int64, 0)
	Type := make([]string, 0)

	for _, val := range arg.Ops {
		if val.ID == "node5_saveas" {
			if !strings.HasPrefix(val.Fop.Result.KodoFile.Key, "encoded") {
				continue
			}
			ResourceKey = append(ResourceKey, val.Fop.Result.KodoFile.Key)
			Quality = append(Quality, 10)
			Type = append(Type, "audio")
			break
		}
	}
	for _, val := range arg.Ops {
		if val.ID == "node6_saveas" {
			ResourceKey = append(ResourceKey, val.Fop.Result.KodoFile.Key)
			Quality = append(Quality, 10)
			Type = append(Type, "video")
			break
		}
	}
	for _, val := range arg.Ops {
		if val.ID == "node7_saveas" {
			ResourceKey = append(ResourceKey, val.Fop.Result.KodoFile.Key)
			Quality = append(Quality, 20)
			Type = append(Type, "video")
			break
		}
	}

	resp, err := s.svr.CallbackTwo(context.Background(), &videoApi.RebackTwoReq{
		//Title:       arg,
		ResourceKey: ResourceKey,
		Type:        Type,
		VideoId:     VideoId,
		Quality:     Quality,
	})
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) FollowList(c *gin.Context) {
	arg := new(struct {
		UserId string `json:"user_id"`
	})
	arg.UserId = c.Query("user_id")
	resp, err := s.svr.FollowList(context.Background(), &baseApi.FollowListReq{
		UserId: arg.UserId,
	})
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) FollowerList(c *gin.Context) {
	arg := new(struct {
		UserId string `json:"user_id"`
	})
	arg.UserId = c.Query("user_id")
	resp, err := s.svr.FollowerList(context.Background(), &baseApi.FollowerListReq{
		UserId: arg.UserId,
	})
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) Follow(c *gin.Context) {
	arg := new(struct {
		UserId     string `json:"user_id"`
		ToUserId   string `json:"to_user_id"`
		ActionType int64  `json:"action_type"`
	})
	if tools.HandleError(c, c.Bind(arg), "") {
		return
	}
	userId, _ := c.Get("user_id")
	arg.UserId = userId.(string)
	resp, err := s.svr.Follow(context.Background(), &baseApi.FollowReq{
		UserId:     arg.UserId,
		ToUserId:   arg.ToUserId,
		ActionType: arg.ActionType,
	})
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) SearchVideoByVideoTag(c *gin.Context) {
	arg := new(struct {
		VideoTag string `json:"video_tag"`
		UserId   string `json:"user_id"`
	})
	arg.VideoTag = c.Query("video_tag")
	userId, _ := c.Get("user_id")
	arg.UserId = userId.(string)
	resp, err := s.svr.SearchVideoByVideoTag(context.Background(), &videoApi.SearchVideoByVideoTagReq{
		VideoTag: arg.VideoTag,
		UserId:   arg.UserId,
	})
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) VideoFeed(c *gin.Context) {
	arg := new(struct {
		UserId string `json:"user_id"`
	})
	userID, _ := c.Get("user_id")
	arg.UserId = userID.(string)
	resp, err := s.svr.VideoFeed(context.Background(), &videoApi.VideoFeedReq{
		LatestTime: 0,
		UserId:     arg.UserId,
	})
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) CollectVideoAction(c *gin.Context) {
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
	resp, err := s.svr.CollectVideoAction(context.Background(), &baseApi.CollectVideoActionReq{
		UserId:     arg.UserId,
		VideoId:    arg.VideoId,
		ActionType: arg.ActionType,
	})
	tools.HandleErrOrResp(c, resp, err)
}

func (s *HttpServer) VideoCollectList(c *gin.Context) {
	arg := new(struct {
		UserId string `json:"user_id"`
	})
	userID, _ := c.Get("user_id")
	arg.UserId = userID.(string)
	resp, err := s.svr.VideoCollectList(context.Background(), &baseApi.VideoCollectListReq{
		UserId: arg.UserId,
	})
	tools.HandleErrOrResp(c, resp, err)
}
