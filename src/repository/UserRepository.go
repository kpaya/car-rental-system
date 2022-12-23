package repository

import (
	"database/sql"
	"errors"

	"github.com/kpaya/car-rental-system/src/entity"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) FindByEmail(email string) entity.User {
	var user entity.User
	result := u.DB.QueryRow("SELECT id, name, email, password, status FROM users WHERE email = $1", email)
	err := result.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status)
	if err != nil {
		return entity.User{}
	}
	return user
}

func (u *UserRepository) Create(user *entity.User) error {
	prep, err := u.DB.Prepare("INSERT INTO users (id, name, email, password, status) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	_, err = prep.Exec(user.ID, user.Name, user.Email, user.Password, user.Status)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) Update(user *entity.User) error {

	foundUser, err := u.FindById(user.ID)
	if err != nil {
		return err
	} else if foundUser.ID == "" {
		return errors.New("this user doesn't exists")
	}
	prep, err := u.DB.Prepare(`UPDATE users SET name = $1, email = $2, password = $3, status = $4 WHERE id = $5`)
	if err != nil {
		return err
	}
	_, err = prep.Exec(user.Name, user.Email, user.Password, user.Status, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) FindById(id string) (entity.User, error) {
	if len(id) == 0 {
		return entity.User{}, errors.New("you must provide a valid id")
	}
	var user entity.User
	prep := u.DB.QueryRow(`SELECT id, name, email, password, status FROM users WHERE id = $1`, id)
	err := prep.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *UserRepository) Delete(id string) error {
	if len(id) == 0 {
		return errors.New("you must provide a valid id")
	}

	u.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return nil
}

func (u *UserRepository) List() ([]entity.User, error) {
	var listUsers []entity.User
	rows, err := u.DB.Query("SELECT id, name, email, password, status FROM users")
	if err != nil {
		return []entity.User{}, err
	}
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status)
		if err != nil {
			return []entity.User{}, err
		}
		listUsers = append(listUsers, user)
	}
	return listUsers, nil
}
