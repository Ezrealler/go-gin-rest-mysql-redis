package configuration

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// ReadConfig default
func ReadConfig() {
	// default ==> setx APP_ENVIRONMENT STAGING
	//Or in main.go init func set os.Setenv("APP_ENVIRONMENT", "STAGING")
	//fmt.Println("os.Getenv", os.Getenv)
	if os.Getenv("APP_ENVIRONMENT") == "STAGING" {
		viper.SetConfigName("properties-staging") //配置文件名
		viper.SetConfigType("yaml")               //配置后缀类型
		viper.AddConfigPath("./resource")         //配置文件所在目录
		//viper.SetConfigFile("/resource/properties-staging")
	} else if os.Getenv("APP_ENVIRONMENT") == "PROD" {
		viper.SetConfigName("properties-prod")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./resource")
		//viper.SetConfigFile("../resource")
	}

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("err viper_config.go")
		fmt.Println(err)
	}
}
