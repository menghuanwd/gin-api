package category

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gin-api/exceptions"
	"gin-api/models"
	"gin-api/params"
	"gin-api/repository"
	"gin-api/repository/category_repo"
	"gin-api/views"
	"gin-api/vos"
)

var (
	err        error
	vo         vos.Category
	category   *models.Category
	paramsBody map[string]interface{}
)

// @Tags 后台 分类
// @Summary 列表
// @Accept application/json
// @Produce application/json
// @Param name query string false "名称"
// @Success 200 {array} vos.Category
// @Router /merchant/categories [get]
func Index(c *gin.Context) {
	var vos []*vos.Category
	customWhere := repository.CondWhereAssemble(c)

	categories, total := category_repo.PageList(customWhere)

	copier.Copy(&vos, &categories)

	views.RenderList(c, vos, total)
}

// @Tags 后台 分类
// @Summary 详细
// @Accept application/json
// @Produce application/json
// @Param id    path    int     true        "id"
// @Success 200 {object} vos.Category
// @Router /merchant/categories/{id} [get]
func Show(c *gin.Context) {
	category = findCategory(c)

	copier.Copy(&vo, &category)

	views.Render(c, vo)
}

// @Tags 后台 分类
// @Summary 创建
// @Accept application/json
// @Produce application/json
// @Param category_model body params.Category true "新增分类"
// @Success 200 {object} vos.Category
// @Router /merchant/categories [post]
func Create(c *gin.Context) {
	var (
		category   models.Category
		paramsBody params.Category
	)

	err = c.ShouldBindJSON(&paramsBody)

	exceptions.Error422(err)

	copier.Copy(&category, &paramsBody)

	err = category.Create()

	exceptions.Error422(err)

	copier.Copy(&vo, &category)

	views.Render(c, vo)
}

// @Tags 后台 分类
// @Summary 更新
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Param category_model body params.Category true "更新分类"
// @Success 200 {object} vos.Category
// @Router /merchant/categories/{id} [put]
func Update(c *gin.Context) {
	err = c.ShouldBindJSON(&paramsBody)

	exceptions.Error422(err)

	category = findCategory(c)

	err = category.Update(paramsBody)

	exceptions.Error422(err)

	err = copier.Copy(&vo, &category)

	exceptions.Error422(err)

	views.Render(c, vo)
}

// @Tags 后台 分类
// @Summary 删除
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Success 204
// @Router /merchant/categories/{id} [delete]
func Destroy(c *gin.Context) {
	category = findCategory(c)

	err := category.Destroy()

	exceptions.Error422(err)

	views.Render204(c)
}


func findCategory(c *gin.Context) (category *models.Category) {
	id := c.MustGet("uintID").(uint)
	category, err = category_repo.Find(id)

	exceptions.Error404(err)

	return
}
