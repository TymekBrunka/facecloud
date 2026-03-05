package endpoints

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	cfg "fcserver/config"
	"log"
	"net/http"
	"os"
)

type reinitREQ struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func ReinitDB(state *cfg.Config, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	var data reinitREQ
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("(REINIT) \x1b[35mfailed to parse json\x1b[0m: ", err)
		http.Error(w, "Wrong request format", http.StatusBadRequest)
		return
	}

	h := sha256.New()
	password := data.Password
	h.Write([]byte(password[4:5] + password + password[2:4]))
	password = hex.EncodeToString(h.Sum(nil))

	if data.Login == "" || data.Password == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		return
	}

	if data.Login != state.REINIT_LOGIN_ || password != state.REINIT_PASSWORD_ {
		http.Error(w, "Wrong credentials", http.StatusBadRequest)
		return
	}

	tx, err := state.Db.Begin()
	if err != nil {
		log.Println("(REINIT) \x1b[31mfailed to create transaction\x1b[0m: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	query_s, err := os.ReadFile("sqlv2.sql")
	if err != nil {
		log.Println("(REINIT) \x1b[35mCant reset database: couldnt find file 'sqlv2.sql'\x1b[0m")
		http.Error(w, "Cant reset database: couldnt find file 'sqlv2.sql'", http.StatusBadRequest)
		return
	}

	_, err1 := tx.Exec(string(query_s))
	_, err2 := tx.Exec(
		`INSERT INTO users (
				first_name,
				last_name,
				email,
				password,
				birth_date,
				last_login,
				bio,
				is_active
    ) VALUES (
				'Zbigniew',
				'Kucharski',
				$1,
				$2,
				'1969-09-11',
				now(),
				'nie pedał, 100% real',
				true
		);`,
		state.SUPERUSER_EMAIL_,
		state.SUPERUSER_PASSWORD_,
	)
	_, err3 := tx.Exec(
		`INSERT INTO users_users_type (
				user_id,
				user_type_id
 	   ) VALUES (
 	   		1,
 	   		4
 	   );
 	   -- ustaw na admina
 	   INSERT INTO users_users_type (
 	   		user_id,
 	   		user_type_id
 	   ) VALUES (
 	   		1,
 	   		3
 	   );`)

	if err := errors.Join(err1, err2, err3); err != nil {
		tx.Rollback()
		log.Println("(REINIT) \x1b[31mfailed to proceed transaction\x1b[0m: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tx.Commit()
}
