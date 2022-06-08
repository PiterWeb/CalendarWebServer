package api

import (
	"github.com/PiterWeb/CalendarWebServer/db"
	"github.com/PiterWeb/CalendarWebServer/models"
)

func LoginUser(user *models.User) error {

	err := db.CheckUser(user)

	if err != nil {
		return err
	}

	return nil

}

func registerUser(user *models.User) error {

	err := db.RegisterUser(user)

	if err != nil {
		return err
	}

	return nil

}

func deleteUser(user *models.User) error {

	err := db.DeleteUser(user)

	if err != nil {
		return err
	}

	return nil

}
