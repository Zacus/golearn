package main

import (
	"context"
	"fmt"
	"golearn/web_app/controller"
	"golearn/web_app/dao/mysql"
	"golearn/web_app/dao/redis"
	"golearn/web_app/logger"
	"golearn/web_app/pkg/snowflake"
	"golearn/web_app/router"
	"golearn/web_app/settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	//1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed,err:%v\n", err)
		return
	}
	fmt.Println("settings success")

	//2.加载日志文件
	if err := logger.Init(settings.Conf.AppConfig.Mode); err != nil {
		fmt.Printf("Log initialization failed", zap.Error(err))
		return
	}
	defer zap.L().Sync()
	zap.L().Info("Log initialization successful")

	//3.初始化数据库
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		zap.L().Error("mysql initialization failed", zap.Error(err))
		return
	}
	defer mysql.Close()
	zap.L().Info("mysql initialization successful")

	//4.初始化redis
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		zap.L().Error("redis initialization failed", zap.Error(err))
		return
	}
	defer redis.Close()
	zap.L().Info("redis initialization successful")

	//初始化雪花算法
	if err := snowflake.Init(settings.Conf.AppConfig.StartTime, int64(settings.Conf.AppConfig.MachineID)); err != nil {
		zap.L().Error("snowflake initialization failed", zap.Error(err))
		return
	}
	zap.L().Info("snowflake initialization successful")
	//初始化gin框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		zap.L().Error("validator initialization failed", zap.Error(err))
		return
	}

	//5.注册路由
	r := router.SetupRouter()
	//6.启动服务(优雅开关机)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.AppConfig.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen:", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")

}
