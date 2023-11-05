package router

import (
	"github.com/Snaptime23/snaptime-server/v2/base/interface/internal/server/http"
	"github.com/Snaptime23/snaptime-server/v2/mw"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateEngine() *gin.Engine {
	if os.Getenv("DEBUG") != "" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	eng := gin.Default()
	eng.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	return eng
}

func InitBaseRouter(engine *gin.RouterGroup, server *http.HttpServer) {
	user := engine.Group("/user")
	{
		user.POST("/register", server.UserRegister)
		user.POST("/login", server.UserLogin)
		user.Use(mw.JwtAuth()).GET("/info", server.UserInfo)
		user.Use(mw.JwtAuth()).GET("/publish/list", server.PublishList)
		user.Use(mw.JwtAuth()).POST("/follow", server.Follow)
		user.Use(mw.JwtAuth()).GET("/follow/list", server.FollowList)
		user.Use(mw.JwtAuth()).GET("/follower/list", server.FollowerList)
	}

	comment := engine.Group("/comment")
	{
		comment.Use(mw.XJwtAuth()).GET("/list", server.CommentList)
		comment.Use(mw.JwtAuth()).POST("/create", server.CreateComment)
		comment.Use(mw.JwtAuth()).POST("/like", server.LikeComment)
	}

	video := engine.Group("/video").Use(mw.JwtAuth())
	{
		video.POST("/like", server.LikeVideoAction)
		video.GET("/like/list", server.VideoLikeList)
		video.GET("/download", server.DownLoadVideo)
		video.GET("/search", server.SearchVideoByVideoTag)

		video.POST("/upload/token", server.UpLoadVideo)
		video.POST("/upload/meta", server.UpLoadVideo)
	}

	callback := engine.Group("/qiniu/callback").Use(mw.QiniuAuth())
	{
		callback.POST("/uploaded", server.CallbackOne)
		callback.POST("/encoded", server.CallbackTwo)
	}
}
