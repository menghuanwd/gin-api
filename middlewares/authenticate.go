package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gin-api/exceptions"
	"gin-api/repository/user_repo"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			exceptions.Error401(errors.New("请求未携带token，无权限访问"))
		}

		user, err := user_repo.FindWechatByAuthToken(token)

		exceptions.Error401(err)

		c.Set("current_user", user)

		c.Next()
	}
}
