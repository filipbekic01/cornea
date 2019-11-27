package services

import "github.com/filipbekic01/cornea/app/models"

type UserService interface {
	GetAll() []models.User
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) GetAll() []models.User {
	return nil
}
