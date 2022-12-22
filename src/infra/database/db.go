package database

import (
	"database/sql"
	"log"
	"os"
)

type Db struct {
	Db *sql.DB
}

func NewDb() *sql.DB {
	conn, err := sql.Open("postgres", os.Getenv("DB_DRIVER"))
	if err != nil {
		log.Panic(err.Error())
	}

	return conn

}
