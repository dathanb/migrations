package models

import "time"

type User struct {
	Id           int    `json:"id"`
	DisplayName  string `json:"display_name"`
	CreationDate Time   `json:"creation_date,string"`
}

func NewUser() User {
	return User{
		Id: 0,
		DisplayName: "",
		CreationDate: ZeroDate(),
	}
}

func ZeroDate() Time {
	time, err := time.Parse("2006-01-02T15:04:05", "1900-01-01T00:00:00.000")
	if err != nil {
		panic("Failed to parse constant time")
	}

	return Time(time)
}

func UserEquals(first User, second User) bool {
	return first.Id == second.Id && first.CreationDate == second.CreationDate && first.DisplayName == second.DisplayName
}

func UserIsBlank(user User) bool {
	return UserEquals(user, NewUser())
}
