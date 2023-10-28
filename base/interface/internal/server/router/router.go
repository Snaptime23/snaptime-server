package router

import (
	"github.com/Snaptime23/snaptime-server/v2/base/interface/internal/server/http"
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

func InitBaseRouter(engine *gin.Engine, server *http.HttpServer) {
	user := engine.Group("/user")
	{
		user.POST("/register", server.UserRegister)
		user.POST("/login", server.UserLogin)
		user.GET("/info", server.UserInfo)
		user.GET("/publish/list", server.PublishList)
	}
}
