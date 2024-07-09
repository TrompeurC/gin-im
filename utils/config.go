package utils

import (
	"fmt"
	"github.com/spf13/viper"
)

type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Addr     string `mapstructure:"addr"`
	Database string `mapstructure:"database"`
}

type Config struct {
	Mysql Mysql `mapstructure:"mysql"`
}

var Conf Config

func InitConfig() {
	viper.SetConfigName("dev")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	err := viper.Unmarshal(&Conf)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
	fmt.Printf("Unmarshal file sucess", "v", Conf)
}
