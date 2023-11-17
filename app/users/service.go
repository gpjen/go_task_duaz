package users

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(input *UserRegister) (User, error)
	Login(input *UserLogin) (UserLoginFormatter, error)
	FindAll() ([]User, error)
	FindByID(id uint) (User, error)
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
		return user, errors.New("email or password is wrong")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userEmail.HashPassword), []byte(input.Password))
	if err != nil {
		return user, errors.New("email or password is wrong")
	}

	user.ID = userEmail.ID
	user.Name = userEmail.Name
	user.Email = userEmail.Email
	user.Address = userEmail.Address

	return user, nil
}

func (s *userService) FindAll() ([]User, error) {

	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *userService) FindByID(id uint) (User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}
