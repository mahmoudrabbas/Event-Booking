package models

import (
	"errors"

	"example.com/events/db"
	"example.com/events/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {

	query := `INSERT INTO users (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return errors.New("Error Preparing The Query.")
	}

	defer stmt.Close()

	hashed, err := utils.HashPassword(u.Password)
	if err != nil {
		return errors.New("Error Hashing The Pass.")

	}

	res, err := stmt.Exec(u.Email, hashed)
	if err != nil {
		return err

	}

	id, err := res.LastInsertId()

	if err != nil {
		return err

	}
	u.Id = id

	return nil

}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	var id int64

	err := row.Scan(&id, &retrievedPassword)

	u.Id = id

	if err != nil {
		return errors.New("Couldnt Found User With This Email.")
	}

	passwordIsValid := utils.CheckPassword(retrievedPassword, u.Password)

	if !passwordIsValid {
		return errors.New("Not Valid Credentials.")
	}

	return nil

}
