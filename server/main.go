package main

import (
	cfg "fcserver/config"
	"fcserver/endpoints"
	"sync"

	// ng "fcserver/netguard"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world!");
	config, err := cfg.Production()
	cfg.Global = cfg.Global_Config_t{Data: config, Lock: sync.RWMutex{}}

	if err != nil {
		log.Fatal("\x1b[31mCannot initialize server\x1b[0m: ", err)
	}

	_res, err := config.Db.Exec("SELECT 1;") // ignore result and check error instead
	if err != nil {
		log.Fatal("\x1b[31m", err, "\x1b[0m")
	}
	_res = _res //ignore unused variable error

	http.HandleFunc("/reinit", endpoints.ReinitDB)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
