package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	AppPort    string
}

func NewConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	return &Config{
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		AppPort:    viper.GetString("APP_PORT"),
	}

}

func (c *Config) BuildPostgresDSN() string {
	dsn := "host=" + c.DBHost +
		" port=" + c.DBPort +
		" user=" + c.DBUser +
		" password=" + c.DBPassword +
		" dbname=" + c.DBName +
		" sslmode=disable TimeZone=Asia/Bangkok"

	return dsn
}
