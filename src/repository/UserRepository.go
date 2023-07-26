package repository

import (
	"database/sql"
	"errors"

	"github.com/kpaya/car-rental-system/src/entity"
	"github.com/kpaya/car-rental-system/src/entity/value_object"
	"golang.org/x/crypto/bcrypt"
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
	prep, err := u.DB.Prepare(
		`INSERT INTO users 
			(id, name, email, password, status)
		VALUES
			($1, $2, $3, $4, $5)
	`)

	if err != nil {
		return err
	}

	defer prep.Close()

	_, err = prep.Exec(user.ID, user.Name, user.Email, user.Password, user.Status)
	if err != nil {
		return err
	}

	prep, err = u.DB.Prepare(
		`INSERT INTO address 
			(id, street_address, city, state, zip_cod, country, user_id)
		VALUES 
			($1,$2,$3,$4,$5,$6,$7)`)
	if err != nil {
		return err
	}
	_, err = prep.Exec(user.Address.ID, user.Address.StreetAddress, user.Address.City, user.Address.State, user.Address.Zipcode, user.Address.Country, user.ID)
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
	rows, err := u.DB.Query(`
	SELECT
		users.id,
		users.name,
		users.email,
		users.type,
		users.password,
		users.status,
		COALESCE(ad.id, CAST('00000000-0000-0000-0000-000000000000' as UUID)) as id,
		COALESCE(ad.street_address, CAST('null' as varchar)) as street_address,
		COALESCE(ad.city, CAST('null' as varchar)) as city,
		COALESCE(ad.state, CAST('null' as varchar))as state,
		COALESCE(ad.zip_cod, CAST('null' as varchar)) as zip_cod,
		COALESCE(ad.country, CAST('null' as varchar)) as country
	FROM users 
		LEFT JOIN address ad ON ad.user_id = users.id
	`)
	if err != nil {
		return []entity.User{}, err
	}
	for rows.Next() {
		var user entity.User
		var address value_object.Address
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Type, &user.Password, &user.Status, &address.ID, &address.StreetAddress, &address.City, &address.State, &address.Zipcode, &address.Country)
		if err != nil {
			return []entity.User{}, err
		}
		user.Address = address
		listUsers = append(listUsers, user)
	}
	return listUsers, nil
}

func (u *UserRepository) FindUserByEmailAndPassword(email string, password string) (entity.User, error) {
	var user entity.User
	row := u.DB.QueryRow("SELECT id, name, email, password, status FROM users WHERE email = $1", email)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status)
	if err != nil {
		return entity.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
