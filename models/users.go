package models

type User struct {
	Id           int    `json:"id"`
	DisplayName  string `json:"display_name"`
	CreationDate Time   `json:"creation_date,string"`
}

