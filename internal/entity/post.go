package entity

import (
	"github.com/uptrace/bun"
	"time"
)

type Post struct {
	bun.BaseModel `bun:"table:posts"`

	Id          int       `json:"id" bun:"id,pk,autoincrement"`
	Description string    `json:"description" bun:"description"`
	UserId      int       `json:"user_id" bun:"user_id"`
	CreatedAt   time.Time `json:"created_at" bun:"created_at"`
}
