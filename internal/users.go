package internal

import (
	"database/sql"
)

type UserModel struct {
	DB *sql.DB
}

func (u UserModel) CreateUser(user User) error {
	stmt := "INSERT INTO users (email, password) VALUES ($1, $2)"
	_, err := u.DB.Exec(stmt, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u UserModel) FindUser(email string) (User, error) {
	stmt := "SELECT email, password FROM users WHERE email = $1"
	row := u.DB.QueryRow(stmt, email)
	var user User
	err := row.Scan(&user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
