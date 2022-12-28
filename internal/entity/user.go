package entity

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	Id       int     `json:"id" bun:"id,pk,autoincrement"`
	Username string  `json:"username" bun:"username"`
	Password string  `json:"password" bun:"password"`
	Avatar   *string `json:"avatar" bun:"avatar"`
}
