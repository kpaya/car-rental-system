package entity

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/kpaya/car-rental-system/src/entity/value_object"
	"golang.org/x/crypto/bcrypt"
)

type AccountStatus string

const (
	Active      AccountStatus = "Active"
	Closed      AccountStatus = "Closed"
	Canceled    AccountStatus = "Canceled"
	Blacklisted AccountStatus = "Blacklisted"
	None        AccountStatus = "None"
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
	ID       string        `json:"id" validate:"uuid4"`
	Name     string        `json:"name" validate:"required"`
	Password string        `json:"-" validate:"required"`
	Status   AccountStatus `json:"status" validate:"required"`
	Email    string        `json:"email" validate:"email,required"`
	Phone    string        `json:"phone" validate:"required"`
	Address  value_object.Location
}

func NewUser(id string, name string, password string, status AccountStatus, email string, phone string) (*User, error) {
	var user = new(User)
	if id == "" {
		id = uuid.NewString()
	}
	if status == "" {
		user.Status = Active
	} else {
		user.Status = status
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
	if err = user.Validate(); err != nil {
		return nil, err
	}
	return user, nil
}

func NewMember(user *User, driverLicenseNumber string, driverLicenseExpire time.Time) *Member {
	var member = new(Member)
	member.User = *user
	member.DriverLicenseExpire = driverLicenseExpire
	member.DriverLicenseNumber = driverLicenseNumber
	return member
}

func NewReceptionist(user *User, dateJoined time.Time) *Receptionist {
	var receptionist = new(Receptionist)
	receptionist.User = *user
	receptionist.DateJoined = dateJoined
	return receptionist
}

func (u *User) Validate() error {
	err := validator.New().Struct(u)
	if err != nil {
		return err
	}
	return nil
}
