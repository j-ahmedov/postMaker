package post_file

import "mime/multipart"

type Filter struct {
	Limit  *int
	Offset *int
}

//type List struct {
//	Id     int    `json:"id"`
//	PostId int    `json:"post_id"`
//	Link   string `json:"link"`
//}
//
//type Detail struct {
//	Id     int    `json:"id"`
//	PostId int    `json:"post_id"`
//	Link   string `json:"link"`
//}

type FileCreate struct {
	PostId int                   `form:"post_id"`
	File   *multipart.FileHeader `form:"file"`
}

type Create struct {
	PostId int    `form:"post_id" json:"post_id"`
	Link   string `form:"link" json:"link"`
}

type Update struct {
	Id     int    `form:"id" json:"id"`
	PostId int    `form:"post_id" json:"post_id"`
	Link   string `form:"link" json:"link"`
}

type FileUpdate struct {
	Id     int                   `form:"id"`
	PostId int                   `form:"post_id"`
	File   *multipart.FileHeader `form:"file"`
}

// Multiple Files

type MultipleList struct {
	Id     int       `json:"id"`
	PostId int       `json:"post_id"`
	Link   *[]string `json:"link"`
}

type MultipleDetail struct {
	Id     int       `json:"id"`
	PostId int       `json:"post_id"`
	Link   *[]string `json:"link"`
}

type MCreate struct {
	PostId int       `form:"post_id" json:"post_id"`
	Link   *[]string `form:"link" json:"link"`
}

type MUpdate struct {
	Id     int       `form:"id" json:"id"`
	PostId int       `form:"post_id" json:"post_id"`
	Link   *[]string `form:"link" json:"link"`
}

type MultipleCreate struct {
	PostId int                     `form:"post_id"`
	Files  []*multipart.FileHeader `form:"files[]"`
}

type MultipleUpdate struct {
	Id     int                     `form:"id"`
	PostId int                     `form:"post_id"`
	Files  []*multipart.FileHeader `form:"files[]"`
}
