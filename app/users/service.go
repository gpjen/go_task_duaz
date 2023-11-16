package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(input *UserRegister) (User, error)
	Login(input *UserLogin) (UserLoginFormatter, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Register(input *UserRegister) (User, error) {
	user := User{}

	duplicateEmail, _ := s.repository.FindBYEmail(input.Email)
	if duplicateEmail.Email != "" {
		return user, errors.New("email already exist")
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Address = input.Address

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}

	user.HashPassword = string(hashPassword)
	newUser, err := s.repository.Register(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *userService) Login(input *UserLogin) (UserLoginFormatter, error) {

	user := UserLoginFormatter{}

	userEmail, err := s.repository.FindBYEmail(input.Email)
	if err != nil {
		return user, err
	}

	user.Name = userEmail.Name
	user.Email = userEmail.Email
	user.Address = userEmail.Address

	return user, nil

}
