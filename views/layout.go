package views

import (
	"github.com/gin-gonic/gin"
	"gin-api/utils"
	"net/http"
)

type Response struct {
	Data interface{} `json:"data"`
}

type MetaPage struct {
	CurrentPage int   `json:"current_page,omitempty"`
	TotalPages  int   `json:"total_pages,default:"0"`
	TotalCounts int64 `json:"total_count,default:"0"`
	Per         int   `json:"per,omitempty"`
}

type ResponseList struct {
	Data interface{} `json:"data"`
	Meta MetaPage    `json:"meta"`
}

func Render204(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNoContent)
}

func Render(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Data: data,
	})
}

func RenderList(c *gin.Context, data interface{}, totalCounts int64) {
	page, per, _ := utils.GetPaginate(c)
	totalPage := utils.GetTotalPages(totalCounts, per)

	c.JSON(http.StatusOK, &ResponseList{
		Data: data,
		Meta: MetaPage{
			page,
			totalPage,
			totalCounts,
			per,
		},
	})
}
