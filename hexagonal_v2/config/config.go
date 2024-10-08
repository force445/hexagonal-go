package config

import (
	"log"

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
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../../")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	config := &Config{
		DBHost:     viper.GetString("POSTGRES_HOST"),
		DBPort:     viper.GetString("POSTGRES_PORT"),
		DBUser:     viper.GetString("POSTGRES_USER"),
		DBPassword: viper.GetString("POSTGRES_PASSWORD"),
		DBName:     viper.GetString("POSTGRES_DB"),
		AppPort:    viper.GetString("APP_PORT"),
	}

	log.Printf("Loaded config: %v\n", config)
	return config

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

func (c *Config) GetAppPort() string {
	return c.AppPort
}
