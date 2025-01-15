package database

import (
	"database/sql"

	"github.com/banisys/user-service/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

var DB *sql.DB

func init() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	DB, err = sql.Open("sqlite3", config.DatabaseUrl)

	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
