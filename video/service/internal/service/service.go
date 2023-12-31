package service

import (
	"context"
	"fmt"
	"math/rand"
	"strings"

	"github.com/Snaptime23/snaptime-server/v2/base/rpc_pb/baseApi"
	"github.com/Snaptime23/snaptime-server/v2/tools/errno"
	"github.com/Snaptime23/snaptime-server/v2/video/rpc_pb/videoApi"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/cache"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/dao"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/dao/model"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/service/downloadToken"
	"github.com/Snaptime23/snaptime-server/v2/video/service/internal/service/uploadToken"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
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
	resp = new(videoApi.VideoFeedResp)
	resp.VideoList = make([]*videoApi.VideoInfo, 0)
	videoIds, err := dao.GetVideoList(ctx)
	if len(videoIds) == 0 {
		return resp, errno.NewErrNo("空空如也")
	}
	cnt := 0
	for i := 0; i < 1000; i++ {
		if cnt > 10 {
			break
		}
		videoId := videoIds[rand.Intn(len(videoIds))]
		video, err := dao.GetVideoByVideoId(ctx, videoId)
		if err != nil {
			continue
		}
		if !(video.UploadState != 0 && video.MetaState != 0) {
			continue
		}
		if video.UploadState == 0 {
			continue
		}
		if video.UploadState == 1 {
			continue
		}
		isEncoding := true
		if video.UploadState == 2 {
			isEncoding = false
		}
		playUrl, err := dao.GetHighUrl(ctx, video.VideoID)
		playUrl = downloadToken.GetToken(playUrl)
		if err != nil && err != gorm.ErrRecordNotFound {
			continue
		}
		if playUrl == "" {
			playUrl = video.PlayUrl
		}
		user, err := s.baseClient.UserInfo(ctx, &baseApi.UserInfoReq{
			UserId: video.CreateUserId,
		})
		if err != nil {
			continue
		}
		hasCollect, err := s.baseClient.HasCollect(ctx, &baseApi.HasCollectReq{
			UserId:  req.UserId,
			VideoId: video.VideoID,
		})
		if err != nil {
			continue
		}
		hasLike, err := s.baseClient.HasLike(ctx, &baseApi.HasLikeReq{
			UserId:  req.UserId,
			VideoId: video.VideoID,
		})
		if err != nil {
			continue
		}
		cnt++
		resp.VideoList = append(resp.VideoList, &videoApi.VideoInfo{
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
				LikeNum:         user.User.FavouriteNum,
				ReceivedLikeNum: user.User.ReceivedLikeNum,
			},
			PlayUrl:       playUrl,
			CoverUrl:      downloadToken.GetToken(video.CoverUrl),
			FavoriteCount: cache.GetVideoLikesCount(video.VideoID),
			CommentCount:  video.CommentCount,
			IsFavorite:    0,
			Title:         video.VideoName,
			CollectCount:  video.CollectCount,
			IsEncoding:    isEncoding,
			HasCollect:    hasCollect.HasCollect,
			HasLike:       hasLike.HasLike,
			Description:   video.Bio,
		})
	}
	return
}

func (s *Service) UploadVideo(ctx context.Context, req *videoApi.UploadVideoReq) (resp *videoApi.UploadVideoResp, err error) {
	resp = new(videoApi.UploadVideoResp)
	videoId := ""
	if req.VideoId != "" {
		videoId = req.VideoId
		err = dao.UpdateVideo(ctx, videoId, &map[string]interface{}{
			"bio":        req.Description,
			"video_name": req.Title,
			"play_url":   "",
			"cover_url":  "",
			"meta_state": 1,
		})
		if err != nil {
			return nil, err
		}
		for _, tag := range req.VideoTag {
			err = dao.CreateVideoTag(ctx, tag, req.VideoId)
			if err != nil {
				continue
			}
		}
		return
	} else {
		videoId = uuid.NewString()
		// return user_upload/{user_uuid}/{video_uuid.file_extension}
		resp.ResourceKey = fmt.Sprintf("user_uploads/%s/%s.%s", req.UserId, videoId, req.FileExtension)
		video := &model.Video{
			VideoID:      videoId,
			Bio:          req.Description,
			VideoName:    req.Title,
			PlayUrl:      downloadToken.GetToken(resp.ResourceKey),
			CoverUrl:     "",
			CreateUserId: req.UserId,
			ResourceKey:  resp.ResourceKey,
		}
		err = dao.CreateVideo(ctx, video)
		if err != nil {
			return nil, err
		}
		resp.Token = uploadToken.GetToken()
		resp.VideoId = videoId
	}
	return
}

func (s *Service) DownLoadVideo(ctx context.Context, req *videoApi.DownloadReq) (resp *videoApi.DownLoadResp, err error) {
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
	if err != nil {
		return
	}
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
		FavoriteCount: cache.GetVideoLikesCount(video.VideoID),
		CommentCount:  video.CommentCount,
		CollectCount:  video.CollectCount,
		IsFavorite:    0,
		Title:         video.VideoName,
		Description:   video.Bio,
	}
	return
}

func (s *Service) CallbackOne(ctx context.Context, req *videoApi.RebackOneReq) (resp *videoApi.RebackOneResp, err error) {
	resp = new(videoApi.RebackOneResp)
	fmt.Println("title = ", req.Title)
	tmp := strings.Split(req.Title, "/")
	if len(tmp) < 2 {
		return
	}
	videoId := tmp[2]
	videoId = strings.ReplaceAll(videoId, ".mp4", "")
	video, err := dao.GetVideoByVideoId(ctx, videoId)
	if err != nil {
		return
	}
	video.UploadState++
	req.CoverUrl = fmt.Sprintf("encoded/screenshots/%s_screenshot.jpg", video.VideoID)
	err = dao.UpdateVideo(ctx, video.VideoID, &map[string]interface{}{
		"upload_state": video.UploadState,
		"cover_url":    req.CoverUrl,
	})
	return
}

func (s *Service) CallbackTwo(ctx context.Context, req *videoApi.RebackTwoReq) (resp *videoApi.RebackTwoResp, err error) {
	resp = new(videoApi.RebackTwoResp)
	video, err := dao.GetVideoByVideoId(ctx, req.VideoId)
	video.UploadState++
	err = dao.UpdateVideo(ctx, video.VideoID, &map[string]interface{}{
		"upload_state": video.UploadState,
	})
	for i := 0; i < len(req.ResourceKey); i++ {
		definition := &model.Definition{
			VideoID:     req.VideoId,
			ResourceKey: req.ResourceKey[i],
			Type:        req.Type[i],
			Quality:     req.Quality[i],
		}
		err = dao.CreateVideoDefinition(ctx, definition)
	}
	return
}

func (s *Service) PublishList(ctx context.Context, req *videoApi.PublishListReq) (resp *videoApi.PublishListResp, err error) {
	resp = new(videoApi.PublishListResp)
	resp.Video = make([]*videoApi.VideoInfo, 0)
	videos, err := dao.GetVideoListByUserId(ctx, req.UserId)
	for _, val := range videos {
		if !(val.UploadState != 0 && val.MetaState != 0) {
			continue
		}
		isEncoding := true
		if val.UploadState == 2 {
			isEncoding = false
		}
		playUrl, err := dao.GetHighUrl(ctx, val.VideoID)
		playUrl = downloadToken.GetToken(playUrl)
		if err != nil && err != gorm.ErrRecordNotFound {
			continue
		}
		if playUrl == "" {
			playUrl = val.PlayUrl
		}
		user, err := s.baseClient.UserInfo(ctx, &baseApi.UserInfoReq{
			UserId: val.CreateUserId,
		})
		if err != nil {
			continue
		}
		hasCollect, err := s.baseClient.HasCollect(ctx, &baseApi.HasCollectReq{
			UserId:  req.UserId,
			VideoId: val.VideoID,
		})
		if err != nil {
			continue
		}
		hasLike, err := s.baseClient.HasLike(ctx, &baseApi.HasLikeReq{
			UserId:  req.UserId,
			VideoId: val.VideoID,
		})
		if err != nil {
			continue
		}
		resp.Video = append(resp.Video, &videoApi.VideoInfo{
			VideoID: val.VideoID,
			Author: &videoApi.UserInfo{
				UserId:          user.User.UserId,
				UserName:        user.User.UserName,
				FollowCount:     user.User.FollowCount,
				FollowerCount:   user.User.FollowerCount,
				IsFollow:        0,
				Avatar:          user.User.Avatar,
				PublishNum:      user.User.PublishNum,
				FavouriteNum:    user.User.FavouriteNum,
				LikeNum:         user.User.FavouriteNum,
				ReceivedLikeNum: user.User.ReceivedLikeNum,
			},
			PlayUrl:       playUrl,
			CoverUrl:      downloadToken.GetToken(val.CoverUrl),
			FavoriteCount: cache.GetVideoLikesCount(val.VideoID),
			CommentCount:  val.CommentCount,
			IsFavorite:    0,
			Title:         val.VideoName,
			IsEncoding:    isEncoding,
			CollectCount:  val.CollectCount,
			HasCollect:    hasCollect.HasCollect,
			HasLike:       hasLike.HasLike,
			Description:   val.Bio,
		})
	}
	return
}

func (s *Service) SearchVideoByVideoTag(ctx context.Context, req *videoApi.SearchVideoByVideoTagReq) (resp *videoApi.SearchVideoByVideoTagResp, err error) {
	resp = new(videoApi.SearchVideoByVideoTagResp)
	resp.Video = make([]*videoApi.VideoInfo, 0)
	videoIds, err := dao.SearchByVideoTag(ctx, req.VideoTag)
	if err != nil {
		return
	}
	if len(videoIds) == 0 {
		return resp, errno.NewErrNo("空空如也")
	}
	cnt := 0
	for i := 0; i < 1000; i++ {
		if cnt > 10 {
			break
		}
		videoId := videoIds[rand.Intn(len(videoIds))]
		video, err := dao.GetVideoByVideoId(ctx, videoId)
		if err != nil {
			continue
		}
		user, err := s.baseClient.UserInfo(ctx, &baseApi.UserInfoReq{
			UserId: video.CreateUserId,
		})
		if err != nil {
			continue
		}
		hasCollect, err := s.baseClient.HasCollect(ctx, &baseApi.HasCollectReq{
			UserId:  req.UserId,
			VideoId: video.VideoID,
		})
		if err != nil {
			continue
		}
		hasLike, err := s.baseClient.HasLike(ctx, &baseApi.HasLikeReq{
			UserId:  req.UserId,
			VideoId: video.VideoID,
		})
		if err != nil {
			continue
		}
		cnt++
		resp.Video = append(resp.Video, &videoApi.VideoInfo{
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
				LikeNum:         user.User.FavouriteNum,
				ReceivedLikeNum: user.User.ReceivedLikeNum,
			},
			PlayUrl:       downloadToken.GetToken(video.ResourceKey),
			CoverUrl:      downloadToken.GetToken(video.CoverUrl),
			FavoriteCount: video.FavouriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    0,
			Title:         video.VideoName,
			CollectCount:  video.CollectCount,
			HasCollect:    hasCollect.HasCollect,
			HasLike:       hasLike.HasLike,
			Description:   video.Bio,
		})
	}
	return
}

func (s *Service) InrcCommentCount(ctx context.Context, req *videoApi.InrcCommentCountReq) (resp *videoApi.InrcCommentCountResp, err error) {
	resp = new(videoApi.InrcCommentCountResp)
	video, err := dao.GetVideoByVideoId(ctx, req.VideoId)
	if err != nil {
		return
	}
	err = dao.UpdateVideo(ctx, req.VideoId, &map[string]interface{}{
		"comment_count": video.CommentCount + req.Count,
	})
	return
}

func (s *Service) UpdateVideo(ctx context.Context, req *videoApi.UpdateVideoReq) (resp *videoApi.UpdateVideoResp, err error) {
	resp = new(videoApi.UpdateVideoResp)
	err = dao.UpdateVideo(ctx, req.Video.VideoID, &map[string]interface{}{
		"play_url":        req.Video.PlayUrl,
		"favourite_count": req.Video.FavoriteCount,
		"cover_url":       req.Video.CoverUrl,
		"comment_count":   req.Video.CommentCount,
		"video_name":      req.Video.Title,
		"collect_count":   req.Video.CollectCount,
	})
	return
}

func (s *Service) IncrFiled(ctx context.Context, req *videoApi.IncrFiledReq) (resp *videoApi.IncrFiledResp, err error) {
	resp = new(videoApi.IncrFiledResp)
	err = cache.IncrVideoField(ctx, req.VideoId, req.Value)
	if err != nil {
		return resp, errno.NewErrNo("update video filed failed")
	}
	return
}
