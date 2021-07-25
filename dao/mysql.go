package dao

import (
	"fmt"
	"gorm.io/gorm/logger"
	"gin-api/exceptions"
	"gin-api/system"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	fmt.Println()
}

func MysqlConnect(config *system.MysqlConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Database,
	)

	newLogger := logger.New(Writer{}, logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Info,
		Colorful:      true,
	})

	MysqlClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: newLogger,
	})

	sqlDB, err := MysqlClient.DB()

	exceptions.Error422(err)

	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return
}
