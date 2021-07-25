package utils

import (
	"github.com/gin-gonic/gin"
)

// GetTotalPages 获取最大数量
func GetTotalPages(totalCounts int64, per int) (totalPages int) {
	totalPages = int(totalCounts) / per
	mod := int(totalCounts) % per
	if mod > 0 {
		totalPages++
	}
	return
}

// RespondWithError ...
func RespondWithError(code int, message string, c *gin.Context) {
	c.AbortWithStatusJSON(code, gin.H{
		"message": message,
	})
}
