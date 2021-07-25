package models

import "gin-api/dao"

type Category struct {
	Base
	Shop     Shop
	Products []Product
	ShopID   uint   `json:"shop_id"`
	Name     string `json:"name" gorm:"index"`
	Image    string `json:"image"`
}

func (category *Category) Create() error {
	return dao.MysqlClient.Create(&category).Error
}

func (category *Category) Update(params map[string]interface{}) error {
	return dao.MysqlClient.Model(&category).Updates(params).Error
}

func (category *Category) Destroy() error {
	return dao.MysqlClient.Delete(Category{}, "id = ?", category.ID).Error
}
