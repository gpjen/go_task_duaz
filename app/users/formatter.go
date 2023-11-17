package users

type UserLoginFormatter struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Role    string `json:"role"`
	Token   string `json:"token"`
}

type UserContex struct {
	ID      uint
	Name    string
	Email   string
	Address string
	Role    string
}
