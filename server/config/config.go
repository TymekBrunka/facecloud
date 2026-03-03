package config

import (
	"database/sql"
	"errors"

	// "log"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type Config struct {
	_REINIT_LOGIN, _REINIT_PASSWORD, _SUPERUSER_EMAIL, _SUPERUSER_PASSWORD, _SUPERUSER_BIRTH_DATE string
	Db                                                                                            *sql.DB
}

// constants from tui/pages/pages.go
const (
	DB = iota
	REINIT_LOGIN
	REINIT_PASSWORD
	SUPERUSER_EMAIL
	SUPERUSER_PASSWORD
	SUPERUSER_BIRTH_DATE

	TEST_DB
	TEST_REINIT_LOGIN
	TEST_REINIT_PASSWORD
	TEST_SUPERUSER_EMAIL
	TEST_SUPERUSER_PASSWORD
	TEST_SUPERUSER_BIRTH_DATE
)

var Keys []string = []string{
	"DB",
	"REINIT_LOGIN",
	"REINIT_PASSWORD",
	"SUPERUSER_EMAIL",
	"SUPERUSER_PASSWORD",
	"SUPERUSER_BIRTH_DATE",

	"TEST_DB",
	"TEST_REINIT_LOGIN",
	"TEST_REINIT_PASSWORD",
	"TEST_SUPERUSER_EMAIL",
	"TEST_SUPERUSER_PASSWORD",
	"TEST_SUPERUSER_BIRTH_DATE",
}

func loadEnv() (map[string]string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		// log.Println("Error loading .env file", err)
		return nil, err
	}

	var env map[string]string

	if env, err = godotenv.Read(); err != nil {
		// log.Println("Error reading the environment variables: %v", err)
		return nil, err
	}

	return env, nil
}

func getEnvKey(env map[string]string, key int16, err *error) string {
	value, exists := env[Keys[key]]
	if exists {
		return value
	} else {
		*err = errors.New(".env missing key " + Keys[key])
		return ""
	}
}

func Production() (Config, error) {
	env, err := loadEnv()

	if err != nil {
		return Config{}, err
	}

	var config Config
	db_key := getEnvKey(env, DB, &err)
	config._REINIT_LOGIN = getEnvKey(env, REINIT_LOGIN, &err)
	config._REINIT_PASSWORD = getEnvKey(env, REINIT_PASSWORD, &err)
	config._SUPERUSER_EMAIL = getEnvKey(env, SUPERUSER_EMAIL, &err)
	config._SUPERUSER_PASSWORD = getEnvKey(env, SUPERUSER_PASSWORD, &err)
	config._SUPERUSER_BIRTH_DATE = getEnvKey(env, SUPERUSER_BIRTH_DATE, &err)

	if err != nil {
		return Config{}, err
	}

	db, err := sql.Open("postgres", db_key)
	if err != nil {
		return Config{}, err
	}

	config.Db = db

	return config, err
}

func Test() (Config, error) {
	env, err := loadEnv()

	if err != nil {
		return Config{}, err
	}

	var config Config
	db_key := getEnvKey(env, DB, &err)
	config._REINIT_LOGIN = getEnvKey(env, REINIT_LOGIN, &err)
	config._REINIT_PASSWORD = getEnvKey(env, REINIT_PASSWORD, &err)
	config._SUPERUSER_EMAIL = getEnvKey(env, SUPERUSER_EMAIL, &err)
	config._SUPERUSER_PASSWORD = getEnvKey(env, SUPERUSER_PASSWORD, &err)
	config._SUPERUSER_BIRTH_DATE = getEnvKey(env, SUPERUSER_BIRTH_DATE, &err)

	if err != nil {
		return Config{}, err
	}

	db, err := sql.Open("postgres", db_key)
	if err != nil {
		return Config{}, err
	}

	config.Db = db

	return config, err
}
