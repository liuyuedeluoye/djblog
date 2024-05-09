package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config = new(AppConfig)

type AppConfig struct {
	Name        string       `mapstructure:"app.name"`
	Mode        string       `mapstructure:"app.mode"`
	Port        string       `mapstructure:"app.port"`
	StartTime   string       `mapstructure:"app.start_time"`
	MachineID   int64        `mapstructure:"app.machine_id"`
	LoggConfig  *LoggConfig  `mapstructure:"log"`
	MysqlConfig *MysqlConfig `mapstructure:"mysql"`
}
type LoggConfig struct {
	Level       string `mapstructure:"level"`
	Filename    string `mapstructure:"filename"`
	Max_size    int    `mapstructure:"max_size"`
	Max_age     int    `mapstructure:"max_age"`
	Max_backups int    `mapstructure:"max_backups"`
}
type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

func Init() (err error) {
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	viper.SetConfigName("config")        // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")          // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath(".")             // 还可以在工作目录中查找配置
	err = viper.ReadInConfig()           // 查找并读取配置文件
	if err != nil {
		// 处理读取配置文件的错误
		return err
	}

	// 手动映射配置值到结构体
	Config.Name = viper.GetString("app.name")
	Config.Mode = viper.GetString("app.mode")
	Config.Port = viper.GetString("app.port")
	Config.StartTime = viper.GetString("app.start_time")
	Config.MachineID = viper.GetInt64("app.machine_id")

	Config.LoggConfig = &LoggConfig{
		Level:       viper.GetString("log.level"),
		Filename:    viper.GetString("log.filename"),
		Max_size:    viper.GetInt("log.max_size"),
		Max_age:     viper.GetInt("log.max_age"),
		Max_backups: viper.GetInt("log.max_backups"),
	}

	Config.MysqlConfig = &MysqlConfig{
		Host:     viper.GetString("mysql.host"),
		Port:     viper.GetString("mysql.port"),
		User:     viper.GetString("mysql.user"),
		Password: viper.GetString("mysql.password"),
		Dbname:   viper.GetString("mysql.dbname"),
	}

	//fmt.Printf("Config: %+v\n", Config)

	viper.WatchConfig()

	//OnConfigChange 是 Viper 提供的一个方法，
	//允许你注册一个回调函数，当配置文件发生更改时，
	//函数将被执行
	// fsnotify 是用于监视文件系统事件的Go库。
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已经修改")
		Config.Name = viper.GetString("app.name")
		Config.Mode = viper.GetString("app.mode")
		Config.Port = viper.GetString("app.port")
		Config.StartTime = viper.GetString("app.start_time")
		Config.MachineID = viper.GetInt64("app.machine_id")

		Config.LoggConfig = &LoggConfig{
			Level:       viper.GetString("log.level"),
			Filename:    viper.GetString("log.filename"),
			Max_size:    viper.GetInt("log.max_size"),
			Max_age:     viper.GetInt("log.max_age"),
			Max_backups: viper.GetInt("log.max_backups"),
		}

		Config.MysqlConfig = &MysqlConfig{
			Host:     viper.GetString("mysql.host"),
			Port:     viper.GetString("mysql.port"),
			User:     viper.GetString("mysql.user"),
			Password: viper.GetString("mysql.password"),
			Dbname:   viper.GetString("mysql.dbname"),
		}
	})
	return nil

}
