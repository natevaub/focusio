package service

import (
	"context"
	"fmt"

	"github.com/natevaub/focus-companion/backend/db/api/dto"
	db "github.com/natevaub/focus-companion/backend/db/generated"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db *db.Queries
}

func NewUserService(db *db.Queries) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(user dto.CreateUserRequest) (*dto.UserResponse, error) {
	// 1. Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// 2. Create user in database with hashed password
	newUser, err := s.db.CreateUser(context.Background(), db.CreateUserParams{
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// 3. Return user response (without password)
	return &dto.UserResponse{
		ID:       int64(newUser.ID),
		Username: newUser.Username,
		Email:    newUser.Email,
	}, nil
}
