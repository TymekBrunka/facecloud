package endpoints

import (
	"encoding/json"
	"log"
	"net/http"
)

type reinitREQ struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func ReinitDB(w http.ResponseWriter, r *http.Request) {
	var data reinitREQ
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Fatal("(REINIT) \x1b[31mfailed to parse json\x1b[0m: ", err)
	}
}
