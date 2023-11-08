package repository

import (
	"crud-go-boilerplate-fiber/app/models/entities"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindByUID(uid string) (*entities.Product, error)
	FindAll() (*[]entities.Product, error)
	Create(db *gorm.DB, input entities.Product) (*entities.Product, error)
	Update(db *gorm.DB, input entities.Product) (*entities.Product, error)
	Delete(db *gorm.DB, id uint64) error
}
type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db}
}
func (r *productRepo) FindAll() (*[]entities.Product, error) {
	var product []entities.Product
	err := r.db.Find(&product).Error
	return &product, err
}
func (r *productRepo) FindByUID(uid string) (*entities.Product, error) {
	var product entities.Product
	err := r.db.Where("product_uid =?", uid).First(&product).Error
	return &product, err
}
func (r *productRepo) Create(db *gorm.DB, input entities.Product) (*entities.Product, error) {
	err := db.Create(&input).Error
	return &input, err
}
func (r *productRepo) Update(db *gorm.DB, input entities.Product) (*entities.Product, error) {
	err := db.Save(&input).Error
	return &input, err
}
func (r *productRepo) Delete(db *gorm.DB, id uint64) error {
	var product entities.Product
	err := db.Where("id =?", id).Delete(&product).Error
	return err
}
