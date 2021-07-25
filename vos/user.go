package vos

type User struct {
	Base
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
