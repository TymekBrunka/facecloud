package main

import (
	cfg "fcserver/config"
	"fcserver/endpoints"

	ng "fcserver/netguard"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world!");
	config, err := cfg.Production()

	if err != nil {
		log.Fatal("\x1b[31mCannot initialize server\x1b[0m: ", err)
	}

	guards := []ng.GuardFunc{
		// ng.G_CSRF_simple,
	}

	ng.Path(&config, guards, "/reinit", endpoints.ReinitDB)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
