package vip

import (
	"fmt"
	"github.com/spf13/viper"
)

// GetIniFile viper 常用读取配置的方法 //
// Get(key string) : interface{}
// GetBool(key string) : bool
// GetFloat64(key string) : float64
// GetInt(key string) : int
// GetIntSlice(key string) : []int
// GetString(key string) : string
// GetStringMap(key string) : map[string]interface{}
// GetStringMapString(key string) : map[string]string
// GetStringSlice(key string) : []string
// GetTime(key string) : time.Time
// GetDuration(key string) : time.Duration
// IsSet(key string) : bool
// AllSettings() : map[string]interface{}

func GetIniFile() {
	config := viper.New()
	config.AddConfigPath("./conf/") // 文件所在目录
	config.SetConfigName("dbuser")  // 文件名
	config.SetConfigType("ini")     // 文件类型

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}

	host := config.GetString("redis.host") // 读取配置
	fmt.Println("viper load ini: ", host)
}

func GetJSONFile() {
	config := viper.New()
	config.AddConfigPath("./conf/")
	config.SetConfigName("host")
	config.SetConfigType("json")

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}

	version := config.GetString("version")
	origin := config.GetString("host.origin")

	fmt.Println(version)
	fmt.Println(origin)

	// 读取到map中
	host := config.GetStringMapString("host")
	fmt.Println(host)
	fmt.Println(host["origin"])
	fmt.Println(host["port"])

	// 所有数据
	allSettings := config.AllSettings()
	fmt.Println(allSettings)
}

func GetYamlFile() {
	config := viper.New()
	config.AddConfigPath("./conf/")
	config.SetConfigName("db")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("找不到配置文件..")
		} else {
			fmt.Println("配置文件出错..")
		}
	}

	host := config.GetString("database.host")
	fmt.Println("viper load yml: ", host)

	allSettings := config.AllSettings()
	fmt.Println(allSettings)

}
