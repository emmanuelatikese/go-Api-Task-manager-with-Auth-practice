package model

type UserModel struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password []byte `json:"password"`
}