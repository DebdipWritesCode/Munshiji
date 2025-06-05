package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver             string `mapstructure:"DB_DRIVER"`
	DBSource             string `mapstructure:"DB_SOURCE"`
	HTTPServerAddress    string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress    string `mapstructure:"GRPC_SERVER_ADDRESS"`
	Environment          string `mapstructure:"ENVIRONMENT"`
	AccessTokenDuration  int64  `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration int64  `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(path string) (Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv() // Override values with system environment variables

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	return config, err
}
