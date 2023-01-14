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
	_, err = conn.Exec("CREATE TABLE IF NOT EXISTS vehicle (id varchar(36), segment_car varchar(100), license_number varchar(50), stock_number varchar(30), passenger_capacity int, barcode varchar(200), has_sunroof bit, status varchar(100), model varchar(300), manufacturing_year int, milage int)")
	if err != nil {
		log.Panic(err.Error())
	}

	return conn
}
