package models

import (
	"github.com/silenceper/wechat/v2/miniprogram/encryptor"
	"gin-api/dao"
)

type User struct {
	Base
	WxaWechat   WxaWechat
	Shop        Shop
	WxaWechatID uint    `json:"wxa_wechat_id"`
	ShopID      uint    `json:"shop_id" gorm:"uniqueIndex:idx_mobile_shop_id"`
	Mobile      string  `json:"mobile" gorm:"size:11;uniqueIndex:idx_mobile_shop_id"`
	UnionID     string  `json:"union_id"`
	MiniOpenID  string  `json:"mini_open_id" gorm:"unique;not null"`
	Nickname    string  `json:"nickname"`
	SessionKey  string  `json:"session_key"`
	Country     string  `json:"country"`
	Province    string  `json:"province"`
	City        string  `json:"city"`
	CountryCode string  `json:"country_code"`
	Language    string  `json:"language"`
	Gender      string  `json:"gender" gorm:"comment:性别"`
	Email       string  `json:"email"`
	Avatar      string  `json:"avatar"`
	Position    string  `json:"position"`
	AuthToken   string  `json:"auth_token"`
	Status      string  `json:"status"`
	Score       int64   `json:"score" gorm:"default:0"`
	Balance     float64 `json:"balance" gorm:"default:0;precision:2`
	Birthday    string  `json:"birthday"`
}

type OrigPlainData struct {
	OpenID          string `json:"openId"`
	UnionID         string `json:"unionId"`
	NickName        string `json:"nickName"`
	Gender          int    `json:"gender"`
	City            string `json:"city"`
	Province        string `json:"province"`
	Country         string `json:"country"`
	AvatarURL       string `json:"avatarUrl"`
	Language        string `json:"language"`
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
}

func UserList(limit, offset int) (result []*User, total int64) {
	db := dao.MysqlClient.Model(&User{})
	db = db.Count(&total)
	db.Offset(offset).Limit(limit).Order("id desc").Find(&result)

	return
}

func (user *User) Create() error {
	res := dao.MysqlClient.Where(User{MiniOpenID: user.MiniOpenID}).
		Assign(User{
			SessionKey:  user.SessionKey,
			WxaWechatID: user.WxaWechatID,
		}).FirstOrCreate(&user)

	return res.Error
}

func (user *User) Upgrade(plainData *encryptor.PlainData) error {
	return dao.MysqlClient.Model(&user).Updates(map[string]interface{}{
		"nickname":          plainData.NickName,
		"avatar_url":        plainData.AvatarURL,
		"union_id":          plainData.UnionID,
		"country":           plainData.Country,
		"province":          plainData.Province,
		"city":              plainData.City,
		"gender":            plainData.Gender,
		"CountryCode":       plainData.CountryCode,
		"language":          plainData.Language,
		"phone_number":      plainData.PhoneNumber,
		"pure_phone_number": plainData.PurePhoneNumber,
	}).Error
}

func (user *User) Update(params map[string]interface{}) error {
	return dao.MysqlClient.Model(&user).Updates(params).Error
}
