package users

type UserLoginFormatter struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Token   string `json:"token"`
}
