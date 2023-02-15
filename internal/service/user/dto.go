package user

import "mime/multipart"

type Filter struct {
	Limit  *int
	Offset *int
}

type List struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type Detail struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type CreateForm struct {
	Username string                `form:"username" json:"username"`
	Password string                `form:"password" json:"password"`
	Avatar   *multipart.FileHeader `form:"avatar" json:"-"`
}

type Create struct {
	Username string  `form:"username" json:"username"`
	Password string  `form:"password" json:"password"`
	Avatar   *string `form:"avatar" json:"avatar"`
}

type Update struct {
	Id       int     `form:"id" json:"id"`
	Username string  `form:"username" json:"username"`
	Password string  `form:"password" json:"password"`
	Avatar   *string `form:"avatar" json:"avatar"`
}

type UpdateForm struct {
	Id       int                   `form:"id" json:"id"`
	Username string                `form:"username" json:"username"`
	Password string                `form:"password" json:"password"`
	Avatar   *multipart.FileHeader `form:"avatar" json:"avatar"`
}

type PostDetail struct {
	Username string  `json:"username"`
	Avatar   *string `json:"avatar"`
}
