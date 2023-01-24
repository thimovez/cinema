package util

import "github.com/spf13/viper"

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	MongoURL     string `mapstructure:"MONGO_URL"`
	RedisURL     string `mapstructure:"REDIS_URL"`
	SecretJwtKey string `mapstructure:"SECRET_JWT_KEY"`
	Scheme       string `mapstructure:"SCHEME"`
	Host         string `mapstructure:"HOST"`
	PathLink     string `mapstructure:"PATH_LINK"`
}

// LoadConfig load environmental variables to the project
func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
