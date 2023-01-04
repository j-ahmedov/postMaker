package entity

import (
	"github.com/uptrace/bun"
)

type PostFile struct {
	bun.BaseModel `bun:"table:post_files"`

	Id     int      `json:"id" bun:"id,pk,autoincrement"`
	PostId int      `json:"post_id" bun:"post_id"`
	Link   []string `json:"link" bun:"link"`
}
