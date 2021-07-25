package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPaginate(c *gin.Context) (page, per, offset int) {
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	per, _ = strconv.Atoi(c.DefaultQuery("per", "10"))
	offset = (page - 1) * per
	return
}
