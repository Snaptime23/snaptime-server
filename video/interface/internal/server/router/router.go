package router

import (
	"github.com/Snaptime23/snaptime-server/v2/video/interface/internal/server/http"
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

func InitVideoRouter(engine *gin.RouterGroup, server *http.HttpServer) {
	video := engine.Group("/video")
	{
		video.POST("/upload_token", server.UpLoadVideo)
		video.GET("/download_token", server.DownLoadVideo)
	}
}
