package post_like

type Filter struct {
	Limit  *int
	Offset *int
}

type List struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	PostId int `json:"post_id"`
}

type Detail struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	PostId int `json:"post_id"`
}

type Create struct {
	UserId int `json:"user_id"`
	PostId int `json:"post_id"`
}

type Update struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	PostId int `json:"post_id"`
}
