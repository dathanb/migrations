package models

type User interface {
	Id() int
	DisplayName() string
}

type dumbUser struct {
	id           int
	display_name string
}

func (usr *dumbUser) Id() int {
	return usr.id
}

func (usr *dumbUser) DisplayName() string {
	return usr.display_name
}

func NewUser(id int, display_name string) User {
	return &dumbUser{
		id:           id,
		display_name: display_name,
	}
}
