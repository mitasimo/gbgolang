package config

import (
	"errors"
	"flag"
	"os"
)

type ServiceConfig struct {
	IPPort, User, Password string
}

func LoadServiceConfig() (*ServiceConfig, error) {

	var cfg ServiceConfig

	flag.StringVar(&cfg.IPPort, "i", ":7319", "Адрес сетевой карты(если несколько):порт, которые будет слушать сервис")
	flag.StringVar(&cfg.User, "u", "", "пользователь сервиса")
	flag.StringVar(&cfg.Password, "p", "", "пароль пользователя сервиса")
	flag.Parse()

	// проверка

	return &cfg, nil
}

type DBConfig struct {
	Host, DB, User, Password string
}

func LoadDBConfig() (*DBConfig, error) {

	var (
		cfg DBConfig
		ok  bool
	)

	cfg.Host, ok = os.LookupEnv("DB_HOST")
	if !ok || cfg.Host == "" {
		return nil, errors.New("не задана переменная окружения DB_HOST")
	}
	cfg.DB, ok = os.LookupEnv("DB_DB")
	if !ok || cfg.DB == "" {
		return nil, errors.New("не задана переменная окружения DB_DB")
	}
	cfg.User, ok = os.LookupEnv("DB_USER")
	if !ok || cfg.User == "" {
		return nil, errors.New("не задана переменная окружения DB_USER")
	}
	cfg.Password, ok = os.LookupEnv("DB_PASSWORD")
	if !ok || cfg.Password == "" {
		return nil, errors.New("не задана переменная окружения DB_PASSWORD")
	}

	return &cfg, nil
}
