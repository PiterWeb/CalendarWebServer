package models

type Date struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

type Event struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Start       Date   `json:"start"`
	End         Date   `json:"end"`
	Important   bool   `json:"important"`
	CalendarId  string `json:"calendarid"`
}

type Calendar struct {
	Username  string   `json:"username"`
	Id string `json:"id"`
}
