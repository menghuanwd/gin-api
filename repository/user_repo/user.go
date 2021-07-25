package user_repo

import (
	"gorm.io/gorm"
	"gin-api/dao"
	"gin-api/models"
	"gin-api/repository"
)

func Find(ID uint) (this models.User, err error) {
	res := dao.MysqlClient.First(&this, ID)

	return this, res.Error
}

func List(cw *repository.CondWhere) (categories []*models.User) {
	filter(cw).Order(cw.Order).Offset(cw.Offset).Limit(cw.Limit).Find(&categories)

	return
}

func Count(cw *repository.CondWhere) (total int64) {
	filter(cw).Count(&total)

	return
}

func PageList(cw *repository.CondWhere) ([]*models.User, int64) {
	return List(cw), Count(cw)
}

func filter(cw *repository.CondWhere) *gorm.DB {
	db := dao.MysqlClient.Model(&models.User{})

	return repository.FilterSql(cw, db)
}

func FindWechatByAuthToken(authToken string) (result *models.User, err error) {
	err = dao.MysqlClient.Where("auth_token = ?", authToken).First(&result).Error

	return result, err
}

func FindWechatByOpenID(OpenID string) (result *models.User, err error) {
	err = dao.MysqlClient.Where("mini_open_id = ?", OpenID).First(&result).Error

	return result, err
}
