package conf

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort string
	AppMode string
	Postgres Postgres
}

type Postgres struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	SSLMode  string
}

func Load() Config {
	godotenv.Load()

	conf := viper.New()
	conf.AutomaticEnv()

	return Config{
		AppPort: conf.GetString("APP_PORT"),
		AppMode: conf.GetString("APP_MODE"),
		Postgres: Postgres{
			User:     conf.GetString("DB_USER"),
			Password: conf.GetString("DB_PASSWORD"),
			Host:     conf.GetString("DB_HOST"),
			Port:     conf.GetString("DB_PORT"),
			Name:     conf.GetString("DB_NAME"),
			SSLMode:  conf.GetString("DB_SSLMODE"),
		},
	}
}
