package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver               string `mapstructure:"DB_DRIVER"`
	DBSource               string `mapstructure:"DB_SOURCE"`
	HTTPServerAddress      string `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress      string `mapstructure:"GRPC_SERVER_ADDRESS"`
	Environment            string `mapstructure:"ENVIRONMENT"`
	AccessTokenDuration    string `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration   string `mapstructure:"REFRESH_TOKEN_DURATION"`
	TokenSymmetricKey      string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	OpenAIAPIKey           string `mapstructure:"OPENAI_API_KEY"`
	RedisAddress           string `mapstructure:"REDIS_ADDRESS"`
	EmailSenderName        string `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress     string `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword    string `mapstructure:"EMAIL_SENDER_PASSWORD"`
	FrontendVerifyEmailURL string `mapstructure:"FRONTEND_EMAIL_VERIFY_URL"`
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
