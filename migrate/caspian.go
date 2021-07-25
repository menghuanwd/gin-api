package main

import (
	"gorm.io/gorm"
	"gin-api/dao"
	sugar "gin-api/log"
	"gin-api/models"
	"gin-api/system"
)

func init() {
	system.LoadConfiguration()
}

func main() {
	dao.Connect()
	defer dao.RedisClient.Close()

	CreateDBTables(dao.MysqlClient)

	sugar.Logger().Info("Finish !!!")
}

func CreateDBTables(orm *gorm.DB) {
	//fmt.Println(orm.Migrator().CurrentDatabase())
	//if !orm.Migrator().HasTable(&models.Category{}) {
	//	orm.Migrator().CreateTable(&models.Category{})
	//} else {
	//	orm.Migrator().DropTable(&models.Category{})
	//}
	//orm.Migrator().DropTable(&models.Category{})
	//
	orm.Migrator().DropTable(
		&models.WxaWechat{},
		&models.Shop{},
		&models.Category{},
		&models.Product{},
		&models.Editor{},
		&models.User{},
	)
	orm.AutoMigrate(
		&models.WxaWechat{},
		&models.Shop{},
		&models.Category{},
		&models.Product{},
		&models.Editor{},
		&models.User{},
	)
}
