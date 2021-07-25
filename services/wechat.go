package services

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"gin-api/system"
)

func InitWechat() *miniprogram.MiniProgram {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &miniConfig.Config{
		AppID:     system.ProjectConfig.Wechat.AppID,
		AppSecret: system.ProjectConfig.Wechat.AppSecret,
		Cache:     memory,
	}
	return wc.GetMiniProgram(cfg)
}

// AuthToken https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxc2fa0eea1a0b2661&secret=881acc2257ab6ea9cc7a88f84387aff8
func AccessToken() (string, error) {
	return InitWechat().GetAuth().GetAccessToken()
}
