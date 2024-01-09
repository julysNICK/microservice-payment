package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DSN string `mapstructure:"DSN"`
	// DBDrive              string        `mapstructure:"DB_DRIVER"`
	// DBSource             string        `mapstructure:"DB_SOURCE"`
	// MigrationUrl         string        `mapstructure:"MIGRATION_URL"`
	// HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	// RedisAddress         string        `mapstructure:"REDIS_ADDRESS"`
	GRPCServerAddress    string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderName      string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress   string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword  string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return

}
