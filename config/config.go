package config

import (
	"github.com/spf13/viper"
	"log"
)

var Config *BCMSConfig

type BCMSConfig struct {
	HttpPort string
	DbDsn    string
}

func ParseConfig(projectName, configFile, env string) {
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("ParseConfigError:", err)
		panic(err)
	}
	Config = new(BCMSConfig)
	err = viper.UnmarshalKey(projectName+"."+env, Config)
	if err != nil {
		log.Println("UnmarshalConfigError:", err)
		panic(err)
	}
}
