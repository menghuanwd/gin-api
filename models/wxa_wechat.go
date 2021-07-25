package models

import "gin-api/dao"

type WxaWechat struct {
	Base
	Users                  []User
	AuthorizerAppid        string `json:"authorizer_appid" gorm:"index"`
	AuthorizerSecret       string `json:"authorizer_secret"`
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
	NickName               string `json:"nick_name"`
	UserName               string `json:"user_name"`
	HeadImg                string `json:"head_img"`
	PrincipalName          string `json:"principal_name"`
	Signature              string `json:"signature"`
	IsVerified             string `json:"is_verified"`
	QrcodeURL              string `json:"qrcode_url"`
}

func FindWxaWechatByAuthorizerAppid(authorizerAppid string) (result User, err error) {
	err = dao.MysqlClient.Where("authorizer_appid = ?", authorizerAppid).First(&result).Error

	return result, err
}
