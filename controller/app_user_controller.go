package controller

import (
	"github.com/WeiDashan/shop-go/global"
	"github.com/WeiDashan/shop-go/service"
	"github.com/WeiDashan/shop-go/service/dto"
	"github.com/WeiDashan/shop-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

type AppUserController struct {
	Service *service.AppUserService
}

func NewAppUserController() AppUserController {
	return AppUserController{
		Service: service.NewAppUserService(),
	}
}

// Login
// @Tag 用户管理
// @Summary 用户登录
// @Description 用户登录详情描述
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {string} string "登录成功"
// @Failure 401 {string} string "登录失败"
// @Router /api/v1/public/user/login [post]
func (m AppUserController) Login(ctx *gin.Context) {
	//ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
	//	"msg": "Login Success",
	//})
	utils.RespSuccess(ctx, "登录成功", nil)
}
func (m AppUserController) Test(ctx *gin.Context) {
	//ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
	//	"msg": "Login Success",
	//})
	utils.RespSuccess(ctx, "test成功", 123)
}
func (m AppUserController) LoginByCode(ctx *gin.Context) {
	var email string
	var code string
	// 获取登录参数：邮箱和验证码
	email = ctx.PostForm("email")
	code = ctx.PostForm("code")
	if email == "" {
		utils.RespError(ctx, "get email error")
		return
	}
	if code == "" {
		utils.RespError(ctx, "get code error")
		return
	}
	// 验证code
	err := m.Service.LoginByCode(email, code)
	if err != nil {
		utils.RespError(ctx, err.Error())
		return
	}
	// 获取AppUser
	appUser, err := m.Service.GetAppUserByEmail(email)
	if err != nil {
		utils.RespError(ctx, err.Error())
		return
	}
	// 生成token
	token, err := utils.GenerateToken(appUser.Id, appUser.NickyName)
	if err != nil {
		utils.RespError(ctx, "token生成失败")
		return
	}
	// 存储token
	err = global.RedisClient.Set(token, appUser.Id, viper.GetDuration("jwt.tokenExpire")*time.Minute)
	if err != nil {
		utils.RespError(ctx, err.Error())
		return
	}
	resultMap := map[string]any{
		"token":   token,
		"appUser": appUser,
	}
	utils.RespSuccess(ctx, "登录成功", resultMap)
}
func (m AppUserController) LoginByPassword(ctx *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO
	errs := ctx.ShouldBind(&iUserLoginDTO)
	if errs != nil {
		utils.RespError(ctx, errs.Error())
		return
	}
	iAppUser, err := m.Service.GetAppUserByEmail(iUserLoginDTO.Email)
	if err != nil {
		utils.RespError(ctx, err.Error())
		return
	}
	token, _ := utils.GenerateToken(iAppUser.Id, iAppUser.NickyName)

	utils.RespSuccess(ctx, "登录成功", gin.H{
		"token":   token,
		"appUser": iAppUser,
	})
}
func (m AppUserController) GetCodeByEmailToLogin(ctx *gin.Context) {
	var email string
	email = ctx.PostForm("email")
	if email == "" {
		utils.RespError(ctx, "get email error")
		return
	}
	err := m.Service.GetCodeByEmailToLogin(email)
	if err != nil {
		utils.RespError(ctx, err.Error())
		return
	}
	utils.RespSuccess(ctx, "正在发送验证码邮件请稍后", nil)
}
