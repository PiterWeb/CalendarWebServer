package db

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"piterdev.com/app/models"
)

func CreateUserCalendarEvent(user User, event models.Event) error {

	userFound := userExists(user.Username)

	if !userFound {
		return errors.New("User not found")
	}

	var calendar models.Calendar

	err := calendarColl.FindOne(ctx, bson.M{"username": user.Username}).Decode(&calendar)

	if err != nil {
		return err
	}

	event.CalendarId = calendar.Id

	_, err = eventColl.InsertOne(ctx, event)

	if err != nil {
		return err
	}

	return nil

}

func getUserEvents(user User) ([]models.Event, error) {

	userFound := userExists(user.Username)

	if !userFound {
		return []models.Event{}, errors.New("User not found")
	}

	if len(user.Password) < 6 {
		return []models.Event{}, errors.New("Username or password are incorrect")
	}

	passwordDecrypted, err := getUserPasswordDecrypted(user.Username)

	if err != nil {
		return []models.Event{}, err
	}

	if user.Password != passwordDecrypted {
		return []models.Event{}, errors.New("Username or password are incorrect")
	}

	var calendar models.Calendar

	err = calendarColl.FindOne(ctx, bson.M{"username": user.Username}).Decode(&calendar)

	if err != nil {
		return []models.Event{}, err
	}

	var events []models.Event

	cur, err := eventColl.Find(ctx, bson.M{"calendarid": calendar.Id})

	if err != nil {
		return []models.Event{}, err
	}

	err = cur.All(ctx, &events)

	if err != nil {
		return []models.Event{}, err
	}

	return events, nil

}

func GetUserCalendarEvents(user User) ([]models.Event, error) {

	events, err := getUserEvents(user)

	if err != nil {
		return []models.Event{}, err
	}

	return events, nil

}

func GetUserCalendarImportantEvents(user User) ([]models.Event, error) {

	events, err := getUserEvents(user)

	if err != nil {
		return []models.Event{}, err
	}

	var importantEvents []models.Event

	for _, event := range events {
		if event.Important == true {
			importantEvents = append(importantEvents, event)
		}
	}

	if len(importantEvents) == 0 {
		return []models.Event{}, errors.New("No important events")
	}

	return importantEvents, nil
}

func GetUserCalendarEventsByDate(user User, date models.Date) ([]models.Event, error) {

	var events []models.Event

	events, err := getUserEvents(user)

	if err != nil {
		return []models.Event{}, err
	}

	for _, event := range events {
		if (event.Start.Day == date.Day && event.Start.Month == date.Month && event.Start.Year == date.Year) || (event.End.Day == date.Day && event.End.Month == date.Month && event.End.Year == date.Year) {
			events = append(events, event)
		}
	}

	if len(events) == 0 {
		return []models.Event{}, errors.New("Event not found")
	}

	return events, nil

}

func GetUserCalendarEventById(user User, id string) (models.Event, error) {

	events, err := getUserEvents(user)

	if err != nil {
		return models.Event{}, err
	}

	for _, event := range events {
		if event.Id == id {
			return event, nil
		}
	}

	return models.Event{}, errors.New("Event not found")

}

func DeleteUserCalendarEventById(user User, id string) error {

	_, err := GetUserCalendarEventById(user, id)

	if err != nil {
		return err
	}

	_, err = eventColl.DeleteOne(ctx, bson.M{"id": id})

	if err != nil {
		return err
	}

	return nil

}

func UpdateUserCalendarEventById(user User, id string, event models.Event) error {

	_, err := GetUserCalendarEventById(user, id)

	if err != nil {
		return err
	}

	_,err = eventColl.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": event})

	if err != nil {
		return err
	}

	return nil

}


