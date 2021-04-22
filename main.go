package main

import (
	"fmt"
	"os"

	cfg "github.com/mitasimo/gbgolang/config"
)

func main() {
	dbConf, err := cfg.LoadDBConfig()
	if err != nil {
		fatal("Ошибка конфигурации базы данных:", err)
	}
	fmt.Println(dbConf)

	svcConf, err := cfg.LoadServiceConfig()
	if err != nil {
		fatal("Ошибка конфигурации сервиса:", err)
	}
	fmt.Println(svcConf)
}

func fatal(messages ...interface{}) {
	fmt.Fprintln(os.Stderr, messages...)
	os.Exit(-1)
}
