package pkg

import (
	"github.com/spf13/viper"
)

type EnvConfig struct {
	WeatherCli_API_KEY string `mapstructure:"API_KEY"`
}

func LoadConfig() (config EnvConfig, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.SetEnvPrefix("WeatherCli")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	return
}
