package params

type WechatUserCode struct {
	AppID string `json:"appid" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

type WechatUserGrantParams struct {
	AppID         string `json:"appid" binding:"required"`
	Iv            string `json:"iv" binding:"required"`
	OpenID        string `json:"open_id" binding:"required"`
	EncryptedData string `json:"encryptedData" binding:"required"`
}

type CodeUnlimitParams struct {
	AppID string `json:"appid" binding:"required"`
	Scene string `json:"scene" binding:"required"`
}

type User struct {
	Mobile      string  `json:"mobile"`
	Nickname    string  `json:"nickname"`
	Country     string  `json:"country"`
	Province    string  `json:"province"`
	City        string  `json:"city"`
	CountryCode string  `json:"country_code"`
	Language    string  `json:"language"`
	Gender      string  `json:"gender"`
	Email       string  `json:"email"`
	Avatar      string  `json:"avatar"`
	Position    string  `json:"position"`
	AuthToken   string  `json:"auth_token"`
	Status      string  `json:"status"`
	Score       int64   `json:"score"`
	Balance     float64 `json:"balance"`
	Birthday    string  `json:"birthday"`
}
