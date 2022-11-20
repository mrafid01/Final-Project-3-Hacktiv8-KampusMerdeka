package service

import (
	"project3/model/entity"
	"project3/model/input"
	"project3/repository"
)

type UserService interface {
	CreateUser(userInput input.UserRegisterInput) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	GetUserByID(id_user int) (entity.User, error)
	UpdateUser(id_user int, input input.UserUpdateInput) (entity.User, error)
	DeleteUser(id_user int) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(input input.UserRegisterInput) (entity.User, error)            {}
func (s *userService) GetUserByEmail(email string) (entity.User, error)                         {}
func (s *userService) GetUserByID(id_user int) (entity.User, error)                             {}
func (s *userService) UpdateUser(id_user int, input input.UserUpdateInput) (entity.User, error) {}
func (s *userService) DeleteUser(id_user int) (entity.User, error)                              {}
