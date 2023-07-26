package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
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

	_, err = conn.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		name varchar(150),
		email varchar(150),
		password varchar(300),
		status varchar(50),
		type char(1) NOT NULL
	)`)

	if err != nil {
		log.Panic(err.Error())
	}

	_, err = conn.Exec(`
	CREATE TABLE IF NOT EXISTS vehicle (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		segment_car varchar(100),
		license_number varchar(50),
		stock_number varchar(30),
		passenger_capacity int,
		barcode varchar(200),
		has_sunroof bit,
		status varchar(100),
		model varchar(300),
		manufacturing_year int,
		milage int
	)`)

	if err != nil {
		log.Panic(err.Error())
	}

	if _, err = conn.Exec(`
		CREATE TABLE IF NOT EXISTS address(
			id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
			street_address VARCHAR(100),
			city VARCHAR(100),
			state VARCHAR(100),
			zip_cod VARCHAR(11),
			country VARCHAR(100),
			user_id uuid REFERENCES users (id)
		)`); err != nil {
		log.Panic(err.Error())
	}

	CreateDefaultAdminAccount(conn)

	return conn
}

func CreateDefaultAdminAccount(db *sql.DB) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), 10)

	if err != nil {
		log.Panic("error to encrypt admin password")
	}

	var adminUser string
	row := db.QueryRow("SELECT id FROM users WHERE email = 'admin'")

	if err = row.Scan(&adminUser); err == nil {
		if adminUser != "" {
			return
		}
	} else {
		_, err = db.Exec(`
			INSERT INTO users (name, email, password, status, type)
			VALUES ('admin', 'admin', $1, '', 'R')
			`, hashedPassword)

		if err != nil {
			log.Panic("error to create admin user")
		}
	}

}
