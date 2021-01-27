package services

import (
	"curso-go/api-echo-orm/entities"
	"curso-go/api-echo-orm/repository"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s UserService) GetByUsername(username string) (*entities.GetUserResponse, error) {
	repo := repository.NewUserRepository(s.db)
	user := repo.GetByUsername(username)

	if user == nil {
		return nil, entities.ErrUserDoesNotExist
	}

	response := entities.GetUserResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	return &response, nil
}
