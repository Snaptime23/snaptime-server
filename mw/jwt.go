package mw

import (
	"fmt"
	"github.com/Snaptime23/snaptime-server/v2/tools"
	"github.com/Snaptime23/snaptime-server/v2/tools/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			tools.SendResp(c, 404, nil, "You don't have permission")
			c.Abort()
			return
		}
		token := strings.Split(auth, " ")[1]

		mc, err := jwt.ParseToken(token)

		fmt.Println("mc =", mc)
		if err != nil {
			c.Abort()
			return
		}

		c.Set("user_id", mc.UserID)
		c.Set("user_name", mc.UserName)

		c.Next()
	}
}
