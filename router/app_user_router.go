package router

import (
	"github.com/WeiDashan/shop-go/controller"
	"github.com/WeiDashan/shop-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitAppUserRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		appUserController := controller.NewAppUserController()

		rgPublicUser := rgPublic.Group("user")
		{
			rgPublicUser.POST("/login", appUserController.Login)
		}

		rgAuthAppUser := rgAuth.Group("app-user")
		rgAuthAppUser.Use(utils.Auth())
		{
			rgAuthAppUser.POST("/loginByPassword", appUserController.LoginByPassword)
			rgAuthAppUser.POST("/loginByCode", appUserController.LoginByCode)
			rgAuthAppUser.POST("/test", appUserController.Test)
			rgAuthAppUser.POST("/getCodeByEmailToLogin", appUserController.GetCodeByEmailToLogin)
			rgAuthAppUser.GET("/:id", func(ctx *gin.Context) {
				ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
					"id":   1,
					"name": "zs",
				})
			})
		}

	})
}
