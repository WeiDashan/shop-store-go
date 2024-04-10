package controller

import (
	"github.com/WeiDashan/shop-go/global"
	"github.com/WeiDashan/shop-go/middleware/rabbitmq"
	"github.com/WeiDashan/shop-go/service"
	"github.com/WeiDashan/shop-go/service/dto"
	"github.com/WeiDashan/shop-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

type TestUserController struct {
	Service *service.TestUserService
}

func NewTestUserController() TestUserController {
	return TestUserController{
		Service: service.NewTestUserService(),
	}
}
func (m TestUserController) AddTestUser(ctx *gin.Context) {
	var iTestUserDTO dto.TestUserAddDTO
	errs := ctx.ShouldBind(&iTestUserDTO)
	if errs != nil {
		utils.RespError(ctx, errs.Error())
		return
	}
	//fmt.Println("----------------------")
	//fmt.Println(iTestUserDTO.RawPassword)

	err := m.Service.AddTestUser(&iTestUserDTO)
	if err != nil {
		utils.RespError(ctx, err.Error())
		return
	}
	//fmt.Println("----------------------")
	//fmt.Println(iTestUserDTO.Password)
	utils.RespSuccess(ctx, "", iTestUserDTO)
}

func (m TestUserController) GetCodeByEmailToLogin(ctx *gin.Context) {
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

	utils.RespSuccess(ctx, "正在生成验证码请稍后", nil)
}
func (m TestUserController) GetCodeByEmailToLogin2(ctx *gin.Context) {
	var email string
	email = ctx.PostForm("email")
	if email == "" {
		utils.RespError(ctx, "get email error")
		return
	}

	emailDTO := rabbitmq.EmailDTO{
		email,
		email,
		email,
	}

	// redis中进行标记
	v, _ := global.RedisClient.Get(emailDTO.To)
	if v != "" {
		_ = global.RedisClient.Delete(emailDTO.To)
	}
	err := global.RedisClient.Set(emailDTO.To, "012448",
		viper.GetDuration("loginCode.codeExpire")*time.Minute)
	if err != nil {
		utils.RespError(ctx, err.Error())
		return
	}

	utils.RespSuccess(ctx, "正在生成验证码请稍后", nil)
}
