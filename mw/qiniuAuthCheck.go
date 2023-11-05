package mw

import (
	"github.com/Snaptime23/snaptime-server/v2/tools"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

func abort(c *gin.Context) {
	tools.SendResp(c, 401, nil, "you don't have permission")
	c.Abort()
}

func QiniuAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" || len(authorization) < 6 || authorization[:6] != "Qiniu " {
			abort(c)
			return
		}

		tmpToken := strings.Split(authorization[6:], ":")
		if len(tmpToken) != 2 || len(tmpToken) < 2 {
			abort(c)
			return
		}
		accessKey := os.Getenv("QINIU_ACCESS_KEY")
		if accessKey != tmpToken[0] {
			abort(c)
			return
		}
		// TODO: 并没有结束检验牛牛的 token
		c.Next()
	}
}
