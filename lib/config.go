package lib

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name     string
	HTTPPort int
}

type DBConfig struct {
	DB_NAME     string
	DB_PORT     string
	DB_HOST     string
	DB_PASSWORD string
	DB_USER     string
}

type TokenConfig struct {
	AccessToken  string
	RefreshToken string
}

type Config struct {
	App      AppConfig
	DB       DBConfig
	Token    TokenConfig
	TimeZone string
}

var Cfg Config

func getAppConfig() AppConfig {
	return AppConfig{
		Name:     getStringOrPanic("APP_NAME"),
		HTTPPort: getIntOrPanic("APP_HTTP_PORT"),
	}
}

func getDBConfig() DBConfig {
	return DBConfig{
		DB_NAME:     getStringOrPanic("DB_NAME"),
		DB_HOST:     getStringOrPanic("DB_HOST"),
		DB_PORT:     getStringOrPanic("DB_PORT"),
		DB_USER:     getStringOrPanic("DB_USER"),
		DB_PASSWORD: getStringOrPanic("DB_PASSWORD"),
	}
}

func getTokenConfig() TokenConfig {
	return TokenConfig{
		AccessToken:  getStringOrPanic("Access_Token"),
		RefreshToken: getStringOrPanic("Refresh_Token"),
	}
}

func LoadConfigByFile(path, fileName, fileType string) Config {
	v := viper.New()
	v.SetConfigName(fileName)
	v.SetConfigType(fileType)
	v.AddConfigPath(path)
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, use automatic environment instead %s\n", err)
	}

	Cfg = Config{
		App:      getAppConfig(),
		DB:       getDBConfig(),
		Token:    getTokenConfig(),
		TimeZone: viper.GetString("TIME_ZONE"),
	}

	return Cfg
}
