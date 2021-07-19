package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"time"
	myredis "xuanxue/dao/redis"
)

// 记录ip中间件
func IpRecorder() func(c *gin.Context) {
	return func(c *gin.Context) {
		ip := c.ClientIP()                                   //获取ip
		today := time.Now().Format("2006-01-02")             // 当天日期
		dateTime := time.Now().Format("2006-01-02 15:04:05") // 格式化时间
		todayIpsKey := fmt.Sprintf("%s%s", myredis.GetRedisKey(myredis.KeyTodayIpZSetPF), today)
		ipDetaillKey := fmt.Sprintf("%s%s", myredis.GetRedisKey(myredis.KeyIpDetailZSetPF), ip)
		pipeline := myredis.Db.Pipeline()
		pipeline.ZIncrBy(todayIpsKey, 1, ip)
		//}
		pipeline.ZAdd(ipDetaillKey, redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: dateTime + ` ` + c.Request.URL.Path + `?` + c.Request.URL.RawQuery,
		})
		_, err := pipeline.Exec()
		if err != nil {
			zap.L().Error("IpRecorder.pipeline.Exec()", zap.Error(err))
		}
	}
}
