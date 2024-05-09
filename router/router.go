package router

import (
	"context"
	"djblog/controller"
	"djblog/logger"
	"djblog/middleware"
	"djblog/settings"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Init() {
	router := gin.New()
	router.Use(logger.GinLogger(), logger.GinRecovery(true))
	//注册路由
	v1 := router.Group("/api/v1")
	//判断如果用户已经登录,就可以访问,否则不可以,可以通过中间件
	v1.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	//获取文章列表
	v1.GET("/articleListUser", controller.GetArticleListUser)
	//获取文章内容
	v1.GET("/articleContent/:id", controller.GetArticle)
	//登录管理员权限
	v1.POST("/logic", controller.LogicHandler)

	//jwt 中间件
	v1.Use(middleware.JWTAuthMiddleware())
	//管理者权限接口
	{
		v1.POST("/Add", controller.AddArticle)
		v1.DELETE("/Delete/:id", controller.DeleteArticle)
		v1.GET("/articleUserRoot", controller.GetArticleListRoot)
	}
	srv := &http.Server{
		Addr:    settings.Config.Port,
		Handler: router,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	log.Println("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")

	////重启
	//go func() {
	//	if err := endless.ListenAndServe(":8080", router); err != nil {
	//		log.Fatalf("listen: %s\n", err)
	//	}
	//
	//	log.Println("Server exiting")
	//}()
	return
}
