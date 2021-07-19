package main

import (
	"fmt"
	"go.uber.org/zap"
	"xuanxue/dao/mysql"
	"xuanxue/dao/redis"
	"xuanxue/logger"
	"xuanxue/router"
	"xuanxue/settings"
)

// @title 玄学宝典
// @version 1.0
// @description 玄学宝典接口文档
// @termOfService https://www.xxg.com
// @Host 192.168.3.180:8082
// @BasePath /api/v1
func main() {

	err := settings.Init()
	if err != nil {
		panic("init config err")
	}
	fmt.Printf("%#v\n", settings.Cfg)

	//2. 初始化日志
	err = logger.Init(settings.Cfg.LoggerConf, settings.Cfg.Mode)
	if err != nil {
		zap.L().Debug("Init log failed", zap.Error(err))
		return
	}
	defer zap.L().Sync() //将缓冲区的日志追加到文件中

	//3. 初始化Mysql
	err = mysql.Init(settings.Cfg.MysqlConf)
	if err != nil {
		zap.L().Debug("Init mysql failed", zap.Error(err))
		return
	}
	defer mysql.Close()

	err = redis.Init(settings.Cfg.RedisConf)
	if err != nil {
		zap.L().Debug("Init mysql failed", zap.Error(err))
		return
	}
	r := router.Setup()
	_ = r.Run(fmt.Sprintf(":%d", settings.Cfg.Port))

}
