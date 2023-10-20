package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ConfigInfo struct {
	Env     string `mapstructure:"env" json:"env"`
	Port    string `mapstructure:"port"`
	AppName string `mapstructure:"app_name"`
	Host    string `mapstructure:"app_url"`
}

var (
	Config ConfigInfo

	configReader = viper.New()
)

func init() {
	loadConfig()
}

func loadConfig() {
	// 设置配置文件
	configReader.SetConfigFile("./config.yaml")

	// 读取配置文件
	err := configReader.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件失败，请检查路径是否正确：", err)
		os.Exit(-1)
	}

	// 解析配置文件
	err = configReader.Unmarshal(&Config)
	if err != nil {
		fmt.Println("解析配置文件失败，请检查配置文件格式是否正确：", err)
		os.Exit(-1)
	}
}
