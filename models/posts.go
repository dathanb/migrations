package models

import "time"

type Post struct {
	Id           int        `json:"id"`
	PostType     int        `json:"post_type_id"`
	UserId       int        `json:"user_id"`
	Body         string     `json:"body"`
	CreationDate time.Time `json:"creation_date"`
}
