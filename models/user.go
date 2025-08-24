package models

import (
	"github.com/Omar-Zarraa/REST-API/db"
	"github.com/Omar-Zarraa/REST-API/utils"
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
