package db

import (
	"errors"

	"github.com/PiterWeb/CalendarWebServer/crypt"
	"github.com/PiterWeb/CalendarWebServer/models"

	"github.com/teris-io/shortid"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterUser(user User) error {

	if len(user.Password) < 6 {
		return errors.New("Password length must be at least 6 characters")
	}

	userFound := userExists(user.Username)

	if userFound {
		return errors.New("User already exists")
	}

	passwordEncrypted, err := crypt.Encrypt(user.Password, crypt.MySecret())

	if err != nil {
		return err
	}

	_, err = userColl.InsertOne(ctx, models.User{
		Username: user.Username,
		Password: passwordEncrypted,
	})

	if err != nil {
		return err
	}

	_, err = calendarColl.InsertOne(ctx, models.Calendar{Id: shortid.MustGenerate(), Username: user.Username})

	return nil

}

func userExists(username string) bool {

	var user models.User

	err := userColl.FindOne(ctx, bson.M{"username": username}).Decode(&user)

	if err != nil {
		return false
	}

	return true

}

func CheckUser(user User) error {

	if len(user.Password) < 6 {
		return errors.New("Username or password are incorrect")
	}

	userFound := userExists(user.Username)

	if !userFound {
		return errors.New("User not found")
	}

	passwordDecrypted, err := getUserPasswordDecrypted(user.Username)

	if err != nil {
		return err
	}

	if user.Password != passwordDecrypted {
		return errors.New("Username or password are incorrect")
	}

	return nil

}

func DeleteUser(user User) error {

	err := CheckUser(user)

	if err != nil {
		return err
	}

	_, err = userColl.DeleteOne(ctx, bson.M{"username": user.Username})

	if err != nil {
		return err
	}

	_, err = calendarColl.DeleteOne(ctx, bson.M{"username": user.Username})

	if err != nil {
		return err
	}

	return nil

}

func getUserPasswordDecrypted(username string) (string, error) {

	var user models.User

	err := userColl.FindOne(ctx, bson.M{"username": username}).Decode(&user)

	if err != nil {
		return "", err
	}

	passwordDecrypted, err := crypt.Decrypt(user.Password, crypt.MySecret())

	if err != nil {
		return "", err
	}

	return passwordDecrypted, nil

}
