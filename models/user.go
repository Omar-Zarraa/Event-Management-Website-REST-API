package models

import (
	"errors"

	"github.com/Omar-Zarraa/Event-Management-Website-REST-API/db"
	"github.com/Omar-Zarraa/Event-Management-Website-REST-API/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

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

func GetAllUsers() ([]User, error) {
	selectQuery := `SELECT * FROM Users`

	rows, err := db.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users = []User{}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
