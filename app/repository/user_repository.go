package repository

import (
	"crud-go-boilerplate-fiber/app/models/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, input entities.User) (*entities.User, error)
	Update(db *gorm.DB, input entities.User) (*entities.User, error)
	FindbyUID(uid string) (*entities.User, error)
	FindbyEmail(email string) (*entities.User, error)
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
func (r *userRepository) Create(db *gorm.DB, input entities.User) (*entities.User, error) {
	err := db.Create(&input).Error
	if err != nil {
		return nil, err
	}
	return &input, nil
}
func (r *userRepository) Update(db *gorm.DB, input entities.User) (*entities.User, error) {
	err := db.Save(&input).Error
	if err != nil {
		return nil, err
	}
	return &input, nil
}
func (r *userRepository) FindbyUID(uid string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("uid =?", uid).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) FindbyEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email =?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
