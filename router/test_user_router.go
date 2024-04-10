package router

import (
	"github.com/WeiDashan/shop-go/controller"
	"github.com/gin-gonic/gin"
)

func InitTestUserRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		testUserController := controller.NewTestUserController()

		//rgPublicUser := rgPublic.Group("user")
		//{
		//	rgPublicUser.POST("/login", userApi.Login)
		//}

		rgAuthTestUser := rgAuth.Group("test-user")
		{
			rgAuthTestUser.POST("/addTestUser", testUserController.AddTestUser)
			rgAuthTestUser.POST("/getCodeByEmailToLogin", testUserController.GetCodeByEmailToLogin)
			rgAuthTestUser.POST("/getCodeByEmailToLogin2", testUserController.GetCodeByEmailToLogin2)
		}

	})
}
