package config

import (
	"os"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан KEY в переменной окружения")
	}
	return &Config{
		Key: key,
	}
}

