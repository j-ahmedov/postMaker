package auth

type TokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
