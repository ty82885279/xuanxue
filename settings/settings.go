package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type AppConf struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Port        int    `mapstructure:"port"`
	Mode        string `mapstructure:"mode"`
	StartTime   string `mapstructure:"start_time"`
	MachineId   int64  `mapstructure:"machine_id"`
	*LoggerConf `mapstructure:"log"`
	*MysqlConf  `mapstructure:"mysql"`
	*RedisConf  `mapstructure:"redis"`
}
type LoggerConf struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}
type MysqlConf struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}
type RedisConf struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

var Cfg = new(AppConf)

func Init() (err error) {
	viper.SetConfigFile("./config/config.yaml")

	err = viper.ReadInConfig()
	// 查找并读取配置文件
	if err != nil { // 处理读取配置文件的错误
		fmt.Errorf("Fatal error config file: %v \n", err)
		return
	}
	err = viper.Unmarshal(Cfg)
	if err != nil {
		fmt.Printf("viper unmarshal failed,err:%v\n", err)
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了")
		err = viper.Unmarshal(Cfg)
		if err != nil {
			fmt.Printf("viper unmarshal failed,err:%v\n", err)
			return
		}
	})
	return
}
