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
	}

	comment := engine.Group("/comment")
	{
		comment.Use(mw.JwtAuth()).POST("/create", server.CreateComment)
		comment.GET("/list", server.CommentList)
	}
}
