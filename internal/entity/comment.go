package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Comment struct {
	bun.BaseModel `bun:"table:comments"`

	Id        int       `json:"id" bun:"id,pk,autoincrement"`
	PostId    int       `json:"post_id" bun:"post_id"`
	UserId    int       `json:"user_id" bun:"user_id"`
	Text      string    `json:"text" bun:"text"`
	CreatedAt time.Time `json:"created_at" bun:"created_at"`
}
