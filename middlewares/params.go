package middlewares

import (
	"github.com/gin-gonic/gin"
	"gin-api/exceptions"
	"gin-api/utils"
)

func ParamsAnalyze() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Param("id") != "" {
			uintID, err := utils.StringToUint(c.Param("id"))

			exceptions.Error422(err)

			c.Set("uintID", uintID)
		}

		c.Next()
	}
}
