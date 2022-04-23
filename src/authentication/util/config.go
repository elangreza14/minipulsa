package util

import (
	"os"
)

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

func LoadConfig(path string) (config Config, err error) {
	res := Config{
		DBDriver: os.Getenv("DB_DRIVER"),
		DBSource: os.Getenv("DB_SOURCE"),
	}
	return res, nil
}
