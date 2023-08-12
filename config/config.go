package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadAppConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME/.mongosteen/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
	// 使用本地 Mailhog 拦截
	viper.Set("email.smtp.host", "localhost")
	viper.Set("email.smtp.port", "1025")
}
