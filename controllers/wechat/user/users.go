package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/silenceper/wechat/v2/miniprogram/qrcode"
	"gin-api/exceptions"
	"gin-api/models"
	"gin-api/params"
	"gin-api/repository/user_repo"
	"gin-api/services"
	"gin-api/views"
	"gin-api/vos"
)


// @Tags 小程序 用户
// @Summary 我
// @Description 我
// @Accept json
// @Produce json
// @Success 200 {object} vos.User
// @Router /wechat/users/me [get]
func Me(c *gin.Context) {
	currentUser := c.MustGet("current_user").(models.User)

	views.Render(c, currentUser)
}

// @Tags 小程序 用户
// @Summary 更新
// @Accept application/json
// @Produce application/json
// @Param id path int true "ID"
// @Param user_model body params.User true "更新"
// @Success 200 {object} vos.User
// @Router /wechat/users/{id} [put]
func Update(c *gin.Context) {
	var (
		err error
		paramsBody map[string]interface{}
		vo     vos.User
	)

	err = c.ShouldBindJSON(&paramsBody)

	exceptions.Error422(err)

	user := findUser(c)

	err = user.Update(paramsBody)

	exceptions.Error422(err)

	err = copier.Copy(&vo, &user)

	exceptions.Error422(err)

	views.Render(c, vo)
}

func findUser(c *gin.Context) (user models.User) {
	id := c.MustGet("uintID").(uint)
	user, err := user_repo.Find(id)

	exceptions.Error404(err)
	return
}


// @Tags 小程序 用户
// @Summary 登录
// @Description 登录
// @Accept json
// @Produce json
// @Param user body params.WechatUserCode true "微信Code"
// @Success 200 {object} vos.User
// @Router /wechat/users [post]
func Create(c *gin.Context) {
	var (
		err  error
		code params.WechatUserCode
	)

	err = c.ShouldBindJSON(&code)

	exceptions.Error422(err)

	app, err := models.FindWxaWechatByAuthorizerAppid(code.AppID)

	exceptions.Error422(err)

	result, err := services.InitWechat().GetAuth().Code2Session(code.Code)

	exceptions.Error422(err)

	var api = models.User{
		MiniOpenID:  result.OpenID,
		SessionKey:  result.SessionKey,
		WxaWechatID: app.ID,
	}

	err = api.Create()

	exceptions.Error422(err)

	views.Render(c, api)
}

// @Summary 授权拿去用户信息
// @Description 授权
// @Tags 小程序 用户
// @Accept json
// @Produce json
// @Param user body params.WechatUserGrantParams true "微信授权"
// @Success 200 {object} vos.User
// @Router /wechat/users/grant [post]
func Grant(c *gin.Context) {
	var api params.WechatUserGrantParams

	err := c.BindJSON(&api)

	exceptions.Error422(err)

	result, err := user_repo.FindWechatByOpenID(api.OpenID)

	exceptions.Error422(err)

	plainData, err := services.InitWechat().GetEncryptor().Decrypt(result.SessionKey, api.EncryptedData, api.Iv)

	exceptions.Error422(err)

	result.Upgrade(plainData)

	views.Render(c, result)
}

// @Summary 生成微信小程序码
// @Description 小程序码
// @Tags 小程序 用户
// @Accept json
// @Produce json
// @Param user body params.CodeUnlimitParams true "微信小程序码"
// @Success 200 {object} vos.User
// @Router /wechat/users/wxacodeunlimit [post]
func CodeUnlimit(c *gin.Context) {
	var coderParams qrcode.QRCoder

	err := c.BindJSON(&coderParams)

	exceptions.Error422(err)

	result, err := services.InitWechat().GetQRCode().GetWXACodeUnlimit(coderParams)

	exceptions.Error422(err)

	views.Render(c, result)
}
