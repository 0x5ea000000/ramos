package services

import (
	"0x5ea000000/ramos/pkg/models"
	"0x5ea000000/ramos/pkg/repositories"
	"0x5ea000000/ramos/pkg/utils"
	"errors"
)

type AuthService interface {
	Register(username, password string) error
	Login(username, password string) (string, error)
}

type authService struct {
	repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) AuthService {
	return &authService{repo: repo}
}

// Register a new user
func (s *authService) Register(username, password string) error {
	// Check if user already exists
	user, err := s.repo.GetByUsername(username)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hash the password and save the user
	hashedPassword, _ := utils.HashPassword(password)
	user = &models.User{
		Username: username,
		Password: hashedPassword,
	}

	if _, err := s.repo.Create(user); err != nil {
		return err
	}

	return nil
}

// Authenticate the user and return a JWT token
func (s *authService) Login(username, password string) (string, error) {
	// Check if user exists
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Check the password
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
