package models

type Post struct {
	Id           int        `json:"id"`
	PostType     int        `json:"post_type_id"`
	UserId       int        `json:"owner_user_id"`
	Body         string     `json:"body"`
	CreationDate Time 		`json:"creation_date"`
}

func NewPost() Post {
	return Post{
		Id: 0,
		PostType: 0,
		UserId: 0,
		Body: "",
		CreationDate: ZeroDate(),
	}
}

func PostEquals(first Post, second Post) bool {
	return first.Id == second.Id &&
		first.PostType == second.PostType &&
		first.UserId == second.UserId &&
		first.Body == second.Body &&
		first.CreationDate == second.CreationDate
}

func PostIsBlank(post Post) bool {
	return PostEquals(post, NewPost())
}
