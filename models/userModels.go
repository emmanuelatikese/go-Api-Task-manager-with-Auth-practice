package routes

type UserModel struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password []byte `json:"password"`
	Confirm_Password []byte `json:"confirmed_password"`

}