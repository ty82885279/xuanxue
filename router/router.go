package router

import (
	"github.com/gin-gonic/gin"
	"xuanxue/controller"
	_ "xuanxue/docs"
	"xuanxue/logger"
	"xuanxue/middleware"
)

func Setup() *gin.Engine {

	r := gin.Default()
	r.Use(middleware.Cors(), logger.GinLogger(), logger.GinRecovery(true), middleware.IpRecorder())
	//r.GET("/doc/*any", gs.WrapHandler(swaggerFiles.Handler))
	//r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	{
		v1.GET("/fortune", controller.FortuneHandler)                // 星座
		v1.GET("/calendar", controller.CalendarHandler)              // 黄历
		v1.GET("/calendar/auspicious", controller.AuspiciousHandler) // 择吉日
		v1.GET("/index", controller.DataHandler)
	}
	return r

}
