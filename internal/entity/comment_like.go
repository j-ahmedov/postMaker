package entity

import "github.com/uptrace/bun"

type CommentLike struct {
	bun.BaseModel `bun:"table:comment_likes"`

	Id        int `json:"id" bun:"id,pk,autoincrement"`
	UserId    int `json:"user_id" bun:"user_id"`
	CommentId int `json:"comment_id" bun:"comment_id"`
}
