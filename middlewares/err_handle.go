package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gin-api/exceptions"
)

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if e, ok := err.(*exceptions.Error); ok {
					c.AbortWithStatusJSON(e.Code, gin.H{
						"message": e.Message,
					})
				} else {
					fmt.Println(err)
					c.JSON(500, "系统错误")
				}
			}
		}()
		c.Next()
	}
}
