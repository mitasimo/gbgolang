package main

import (
	"fmt"
	"log"

	cfg "github.com/mitasimo/gbgolang/config"
)

func main() {
	var err error

	sCfg, err := cfg.LoadServiceConfig("service.json")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(sCfg)

	dbCfg, err := cfg.LoadDBConfig("db.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(dbCfg)
}
