package post

import (
	"postMaker/internal/service/comment"
	"postMaker/internal/service/user"
)

type Filter struct {
	Limit  *int
	Offset *int
}

type List struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
	CreatedAt   string `json:"created_at"`
}

type Detail struct {
	Id          int                  `json:"id"`
	Description string               `json:"description"`
	CreatedAt   string               `json:"created_at"`
	Files       []string             `json:"files"`
	User        user.PostDetail      `json:"user"`
	Comments    []comment.PostDetail `json:"comments"`
}

type Create struct {
	Description string `json:"description"`
	UserId      int    `json:"-"`
}

type Update struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}
