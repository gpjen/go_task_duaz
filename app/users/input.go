package users

type UserRegister struct {
	Name            string `json:"name" form:"name" validate:"required,min=3,max=100"`
	Email           string `json:"email" form:"email" validate:"required,email"`
	Address         string `json:"address" form:"address" validate:"required,min=3,max=100"`
	Password        string `json:"password" form:"password" validate:"required,min=6,max=100"`
	ConfirmPassword string `json:"confirmPassword" form:"confirm_password" validate:"required,eqfield=Password"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6,max=100"`
}

type UserUpdate struct {
	Name    string `json:"name" form:"name" validate:"required,min=3,max=100"`
	Email   string `json:"email" form:"email" validate:"required,email"`
	Address string `json:"address" form:"address" validate:"required,min=3,max=100"`
}

type UserUpdatePassword struct {
	OldPassword     string `json:"oldPassword" form:"old_password" validate:"required,min=6,max=100"`
	NewPassword     string `json:"newPassword" form:"new_password" validate:"required,min=6,max=100"`
	ConfirmPassword string `json:"confirmPassword" form:"confirm_password" validate:"required,min=6,max=100"`
}
