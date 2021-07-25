package category_repo

import (
	"gorm.io/gorm"
	"gin-api/dao"
	"gin-api/models"
	"gin-api/repository"
)

func Find(ID uint) (category *models.Category, err error) {
	res := dao.MysqlClient.First(&category, ID)

	return category, res.Error
}

func List(cw *repository.CondWhere) (categories []*models.Category) {
	filter(cw).Order(cw.Order).Offset(cw.Offset).Limit(cw.Limit).Find(&categories)

	return
}

func Count(cw *repository.CondWhere) (total int64) {
	filter(cw).Count(&total)

	return
}

func PageList(cw *repository.CondWhere) ([]*models.Category, int64) {
	return List(cw), Count(cw)
}

func filter(cw *repository.CondWhere) *gorm.DB {
	db := dao.MysqlClient.Model(&models.Category{})

	return repository.FilterSql(cw, db)
}
