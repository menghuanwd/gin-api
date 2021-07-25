package system

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"os"
)

type Config struct {
	Redis  *RedisConfig  `toml:"redis"`
	Mysql  *MysqlConfig  `toml:"mysql"`
	Wechat *WechatConfig `toml:"wechat"`
}

type RedisConfig struct {
	Host     string `toml:"host"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

type MysqlConfig struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Port     string `toml:"port"`
	Database string `toml:"database"`
}

type WechatConfig struct {
	AppID     string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
}

var (
	ProjectConfig *Config
	path          string
)

func LoadConfiguration() error {
	gin.SetMode(os.Getenv("GIN_MODE"))

	if gin.Mode() == gin.ReleaseMode {
		path = "config/production.toml"
	} else {
		path = "config/development.toml"
	}

	_, err := toml.DecodeFile(path, &ProjectConfig)

	return err
}
