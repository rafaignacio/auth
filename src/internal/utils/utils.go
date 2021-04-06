package utils

import "github.com/spf13/viper"

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DBServer      string `mapstructure:"DB_SERVER"`
}

func LoadConfig(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	return
}
