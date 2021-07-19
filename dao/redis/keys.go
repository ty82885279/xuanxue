package redis

const (
	Prefix            = "xuanxue:" // 项目前缀
	KeyTodayIpZSetPF  = "ips:"     // set:统计当天访问的所有ip 参数：当天日期 2020-11-11
	KeyIpDetailZSetPF = "ip:"      // hash:统计每个访问时间，访问路径 参数：ip
)

func GetRedisKey(key string) string {

	return Prefix + key
}
