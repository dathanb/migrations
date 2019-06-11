package models

type Post struct {
	Id           int        `json:"id"`
	PostType     int        `json:"post_type_id"`
	UserId       int        `json:"owner_user_id"`
	Body         string     `json:"body"`
	CreationDate Time 		`json:"creation_date"`
}

