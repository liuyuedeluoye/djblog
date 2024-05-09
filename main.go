package main

import (
	"djblog/dao"
	"djblog/logger"
	"djblog/pkg/snowflake"
	"djblog/router"
	"djblog/settings"
	"fmt"
	"go.uber.org/zap"
)

//本项目是一个dj博客项目

//项目步骤:

// 1.加载配置
func main() {
	//加载viper
	if err := settings.Init(); err != nil {
		fmt.Printf("settings failed,err:%v\n", err)
		return
	}
	fmt.Printf("setting 成功\n")
	//logger
	if err := logger.Init(settings.Config.LoggConfig, settings.Config.Mode); err != nil {
		zap.L().Error("logger init failed", zap.Error(err))
	}
	zap.L().Info("logger success")
	//在Go语言中，日志库通常会使用缓冲区来提高性能，当缓冲区满了或者达到一定的时间间隔时，才会将日志写入到输出目标中。Sync()方法的作用就是手动触发这个过程，将缓冲区的日志写入到输出目标中。
	//写入缓冲区的日志
	defer zap.L().Sync()
	//mysql init
	if err := dao.Init(settings.Config.MysqlConfig); err != nil {
		zap.L().Error("mysql init failed", zap.Error(err))
	}
	zap.L().Info("mysql success")
	defer dao.Close()

	//初始化雪花算法
	if err := snowflake.Init(settings.Config.StartTime, settings.Config.MachineID); err != nil {
		fmt.Printf("init failed,err:%v", err)
		return
	}
	fmt.Println(snowflake.GenID())
	router.Init()
}

//数据库的建立

//日志建立

//配置状态码

//接口配置
