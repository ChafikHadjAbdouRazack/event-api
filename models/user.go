package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	statement, error := db.DB.Prepare(query)

	if error != nil {
		return error
	}

	defer statement.Close()
	hashedPassword, error := utils.HashPassword(user.Password)
	if error != nil {
		return error
	}
	result, error := statement.Exec(user.Email, hashedPassword)

	if error != nil {
		return error
	}

	id, error := result.LastInsertId()

	if error != nil {
		return error
	}

	user.ID = id

	return nil
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id,password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)
	var retrivedPassword string
	error := row.Scan(user.ID, &retrivedPassword)

	if error != nil {
		return error
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrivedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
