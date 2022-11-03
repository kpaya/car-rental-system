package entity

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/kpaya/car-rental-system/src/entity/value_object"
	"golang.org/x/crypto/bcrypt"
)

type Member struct {
	User
	DriverLicenseNumber string
	DriverLicenseExpire time.Time
}

type Receptionist struct {
	User
	DateJoined time.Time
}

type User struct {
	ID       string
	Name     string
	Password string
	Email    string
	Phone    string
	Address  value_object.Location
}

func NewUser(id string, name string, password string, email string, phone string) *User {
	var user = new(User)
	if id == "" {
		id = uuid.NewString()
	}
	user.ID = id
	user.Name = name
	user.Email = email
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Panic("error to encrypt password from user")
	}
	user.Password = string(hashedPassword)
	user.Phone = phone
	return user
}

func (u *User) Validate() {
	err := validator.New().Struct(u)
	if err != nil {
		log.Panic(err)
	}
}
