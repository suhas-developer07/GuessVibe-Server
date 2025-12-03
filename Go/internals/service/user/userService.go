package services

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	models "github.com/suhas-developer07/GuessVibe-Server/internals/models/User_model"
)

type UserService struct {
	userRepo models.UserRepo
	validate *validator.Validate
}

func NewUserService(userRepo models.UserRepo) *UserService {
	v := validator.New()
	return &UserService{
		userRepo: userRepo,
		validate: v,
	}
}
func (s *UserService) RegisterUser(req models.User) (string, error) {
	if err := s.validate.Struct(req); err != nil {
		return "", fmt.Errorf("validation error: %w", err)
	}

	id, err := s.userRepo.RegisterUser(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *UserService) LoginUser(email, password string) (string, error) {
	return s.userRepo.LoginUser(email, password)
}

func (s *UserService) LogoutUser(userID, token string) error {
	if userID == "" {
		return fmt.Errorf("invalid user id")
	}

	return s.userRepo.LogoutUser(userID, token)
}
                         