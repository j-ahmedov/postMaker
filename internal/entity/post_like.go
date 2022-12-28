package entity

import "github.com/uptrace/bun"

type PostLike struct {
	bun.BaseModel `bun:"table:post_likes"`

	Id     int `json:"id" bun:"id,pk,autoincrement"`
	UserId int `json:"user_id" bun:"user_id"`
	PostId int `json:"post_id" bun:"post_id"`
}
