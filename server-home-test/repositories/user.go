package repositories

import (
	"github.com/octadsp/server-home-test/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}
