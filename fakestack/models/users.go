package models

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

func UserEquals(first User, second User) bool {
	return first.Id == second.Id &&
		first.CreationDate == second.CreationDate &&
		first.DisplayName == second.DisplayName
}

func UserIsBlank(user User) bool {
	return UserEquals(user, NewUser())
}
