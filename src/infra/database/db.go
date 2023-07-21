package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type Db struct {
	Db *sql.DB
}

func NewDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%v user=%s "+
		"password=%s dbname=%s sslmode=disable", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB_NAME"))
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err.Error())
	}
<<<<<<< HEAD
	_, err = conn.Exec("CREATE TABLE IF NOT EXISTS users (id UUID DEFAULT gen_random_uuid() PRIMARY KEY, name varchar(150), email varchar(150), password varchar(300), status varchar(50))")
	if err != nil {
		log.Panic(err.Error())
	}
	_, err = conn.Exec("CREATE TABLE IF NOT EXISTS vehicle (id UUID DEFAULT gen_random_uuid() PRIMARY KEY, segment_car varchar(100), license_number varchar(50), stock_number varchar(30), passenger_capacity int, barcode varchar(200), has_sunroof bit, status varchar(100), model varchar(300), manufacturing_year int, milage int)")
=======
	_, err = conn.Exec("CREATE TABLE IF NOT EXISTS users (id varchar(36), name varchar(150), email varchar(150), password varchar(300), status varchar(50))")
	if err != nil {
		log.Panic(err.Error())
	}
	_, err = conn.Exec("CREATE TABLE IF NOT EXISTS vehicle (id varchar(36), segment_car varchar(100), license_number varchar(50), stock_number varchar(30), passenger_capacity int, barcode varchar(200), has_sunroof bit, status varchar(100), model varchar(300), manufacturing_year int, milage int)")
>>>>>>> 42239aee2d434c956ec3b916ea87720f6ec82609
	if err != nil {
		log.Panic(err.Error())
	}

	return conn
}
