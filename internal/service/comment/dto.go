package comment

import "postMaker/internal/service/user"

type Filter struct {
	Limit  *int
	Offset *int
}

type List struct {
	Id        int    `json:"id"`
	PostId    int    `json:"post_id"`
	UserId    int    `json:"user_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

type Detail struct {
	Id        int    `json:"id"`
	PostId    int    `json:"post_id"`
	UserId    int    `json:"user_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}

type Create struct {
	PostId int    `json:"post_id"`
	UserId int    `json:"user_id"`
	Text   string `json:"text"`
}

type Update struct {
	Id     int    `json:"id"`
	PostId int    `json:"post_id"`
	UserId int    `json:"user_id"`
	Text   string `json:"text"`
}

type PostDetail struct {
	Id        int             `json:"id"`
	Text      string          `json:"text"`
	User      user.PostDetail `json:"user"`
	CreatedAt string          `json:"created_at"`
}
