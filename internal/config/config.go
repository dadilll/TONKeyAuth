package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPServerPort int    `env:"HTTP_SERVER_PORT" env-default:"8080"`
	PrivateKeyPath string `env:"PRIVATE_KEY_PATH" env-default:"key/private.pem"`
	PublicKeyPath  string `env:"PUBLIC_KEY_PATH" env-default:"key/public.pem"`
	Issuer         string `env:"ISSUER" env-default:"TON-OAUTH"`
	KeyName        string `env:"KEY_NAME" env-default:"main-key"`
}

func New() *Config {
	cfg := Config{}
	err := cleanenv.ReadConfig("conf/conf.env", &cfg)
	if err != nil {
		fmt.Printf("Ошибка чтения конфигурации: %v\n", err)
		return nil
	}

	fmt.Printf("Загружена конфигурация: %+v\n", cfg)
	return &cfg
}
