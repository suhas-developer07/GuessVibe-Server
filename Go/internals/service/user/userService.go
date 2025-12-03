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
func (s *UserService) RegisterUser(req models.User) (int64, error) {
	err := s.validate.Struct(req)
	if err != nil {
		return 0, fmt.Errorf("validation error: %w", err)
	}
	id, err := s.userRepo.RegisterUser(req)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (s *UserService) LoginUser(Email, password string) (string, error) {
	token, err := s.userRepo.LoginUser(Email, password)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (s *UserService) LogoutUser(UserID, token string) error {
	if UserID == "" {
		return fmt.Errorf("invalid user id")
	}
	err := s.userRepo.LogoutUser(UserID, "")
	if err != nil {
		return err
	}
	return nil
}
