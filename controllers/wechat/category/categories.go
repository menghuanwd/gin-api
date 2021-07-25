package category

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gin-api/repository"
	"gin-api/repository/category_repo"
	"gin-api/views"
	"gin-api/vos"
)

// @Tags 小程序 分类
// @Summary 列表
// @Description 所有分类列表-分页
// @Accept application/json
// @Produce application/json
// @Param name    query    string     false        "名称"
// @Success 200 {array} vos.Category
// @Router /wechat/categories [get]
func Index(c *gin.Context) {
	var vos []vos.Category

	customWhere := repository.CondWhereAssemble(c)

	categories, total := category_repo.PageList(customWhere)

	copier.Copy(&vos, &categories)

	views.RenderList(c, vos, total)
}
