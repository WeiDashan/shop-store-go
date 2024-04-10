package router

import (
	"context"
	"fmt"
	_ "github.com/WeiDashan/shop-go/docs"
	"github.com/WeiDashan/shop-go/global"
	"github.com/WeiDashan/shop-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type IFnRegistRoute = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRoutes []IFnRegistRoute
)

// RegistRoute 注册路由回调函数
func RegistRoute(fn IFnRegistRoute) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

// InitRouter 初始化系统路由
func InitRouter() {

	// 监听ctrl+c，应用推出信号的上下文
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	// 初始化gin框架，并注册相关路由
	r := gin.Default()
	r.Use(utils.Cors())
	//rgPublic := r.Group("/api/v1/public")
	//rgAuth := r.Group("/api/v1")
	rgPublic := r.Group("/public")
	rgAuth := r.Group("")

	//rgAuth.Use(utils.Auth())

	// 初始基础平台的路由
	InitBasePlatformRoutes()

	// 开始注册系统各模块对应的路由信息
	for _, fnRegistRoute := range gfnRoutes {
		fnRegistRoute(rgPublic, rgAuth)
	}

	// 集成swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 从配置文件中读取并配置web服务设置
	port := viper.GetString("server.port")
	if port == "" {
		port = "8080"
	}

	// 创建web server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	// 启动一个goroutine来开启web服务，避免主线程的信号监听被阻塞
	go func() {
		global.Logger.Info(fmt.Sprintf("Start Listen: %s", port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("Start Server Error: %s", err.Error()))
			return
		}
	}()

	// 等待停止服务的信号被触发
	<-ctx.Done()

	// 关闭Server，5秒内未完成清理动作会直接退出应用
	ctx, cancelShutDown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutDown()
	if err := server.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("Stop Server Error: %s", err.Error()))
		return
	}
	global.Logger.Info("Stop Server Success")
}
func InitBasePlatformRoutes() {
	//InitUserRoutes()
	InitAppUserRoutes()
	InitTestUserRoutes()
}
