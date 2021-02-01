package repository

import (
	"curso-go/api-echo-orm/entities"
	"gorm.io/gorm"
)

type IUser interface {
	GetByUsername(username string) *entities.User
	Create(user *entities.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r UserRepository) GetByUsername(username string) *entities.User {
	user := &entities.User{}
	r.db.Where("username = ?", username).First(user)

	if user.Username == "" {
		return nil
	}

	return user
}

func (repo *UserRepository) Create(user *entities.User) error {
	result := repo.db.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}


