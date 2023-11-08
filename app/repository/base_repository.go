package repository

import "gorm.io/gorm"

type BaseRepository interface {
	GetDB() *gorm.DB
	GetBegin() *gorm.DB
}
type baseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return &baseRepository{db}
}

func (br *baseRepository) GetDB() *gorm.DB {
	return br.db
}

func (br *baseRepository) GetBegin() *gorm.DB {
	return br.GetDB().Begin()
}
