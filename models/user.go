package models

import (
	"errors"

	"github.com/Omar-Zarraa/Event-Management-Website-REST-API/db"
	"github.com/Omar-Zarraa/Event-Management-Website-REST-API/utils"
)

//User specifies the structure of the users to be used in the database.
type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

//Save inserts into the Users table the data of the user used to call this method, possibly returns an error.
func (user User) Save() error {
	query := "INSERT INTO Users(Email,Password) VALUES (?,?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	user.ID = userID

	return err
}

//ValidateUser checks if the user data in the user that was used to call this method are valid and are found in the database, possibly returns an error.
func (user *User) ValidateUser() error {
	query := "SELECT ID, Password FROM Users WHERE Email=?"
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPass string
	err := row.Scan(&user.ID, &retrievedPass)
	if err != nil {
		return errors.New("Credentials invalid")
	}

	if !utils.CheckPasswordHash(user.Password, retrievedPass) {
		return errors.New("Credentials invalid")
	}

	return nil
}
