package main

import (
	cfg "fcserver/config"
	"fmt"
	"log"
)

func main() {
	fmt.Println("hello world!");
	config, err := cfg.Production()

	if err != nil {
		log.Fatal("\x1b[31mCannot initialize server\x1b[0m: ", err)
	}

	_res, err := config.Db.Exec("SELECT 1;") // ignore result and check error instead
	if err != nil {
		log.Fatal("\x1b[31m", err, "\x1b[0m")
	}
	_res = _res //ignore unused variable error
}
