package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConfigInfo struct {
	Host    string      `mapstructure:"host"`
	Port    string      `mapstructure:"port"`
	MongoDB MongoDBInfo `mapstructure:"mongodb"`
}

type MongoDBInfo struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Database string `mapstructure:"db"`
}

var (
	Config ConfigInfo

	configReader = viper.New()
	Log          = logrus.New()
)

func init() {
	loadConfig()
	initLog()
}

func loadConfig() {
	// 设置配置文件
	configReader.SetConfigFile("./config.yaml")

	// 读取配置文件
	err := configReader.ReadInConfig()
	if err != nil {
		Log.Error("读取配置文件失败，请检查路径是否正确：", err)
		os.Exit(-1)
	}

	// 解析配置文件
	err = configReader.Unmarshal(&Config)
	if err != nil {
		Log.Error("解析配置文件失败，请检查配置文件格式是否正确：", err)
		os.Exit(-1)
	}
}

func initLog() {
	Log.SetLevel(logrus.DebugLevel)
}
