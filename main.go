package main

import (
	"fmt"
	"log"

	cfg "github.com/mitasimo/gbgolang/config"
)

func main() {
	dbConf, err := cfg.LoadDBConfig()
	if err != nil {
		log.Fatalf("Ошибка конфигурации базы данных: %v\n", err)
	}
	fmt.Println(dbConf)

	svcConf, err := cfg.LoadServiceConfig()
	if err != nil {
		log.Fatalf("Ошибка конфигурации сервиса: %v\n", err)
	}
	fmt.Println(svcConf)
}
