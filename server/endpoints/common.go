package endpoints

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	cfg "fcserver/config"
	"log"
	"net/http"
)

type reinitREQ struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func ReinitDB(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	var data reinitREQ
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("(REINIT) \x1b[31mfailed to parse json\x1b[0m: ", err)
		http.Error(w, "Wrong request format", http.StatusBadRequest)
		return
	}

	h := sha256.New()
	password := data.Password
	h.Write([]byte(password[4:5] + password + password[2:4]))
	password = hex.EncodeToString(h.Sum(nil))

	if data.Login == "" || data.Password == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
	}

	if data.Login != cfg.Global.Data.REINIT_LOGIN_ || data.Password != cfg.Global.Data.REINIT_PASSWORD_ {
		http.Error(w, "Wrong credentials", http.StatusBadRequest)
	}

	tx, err := cfg.Global.Data.Db.Begin()
	if err != nil {
		log.Println("(REINIT) \x1b[31mfailed to create transaction\x1b[0m: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tx = tx
}
