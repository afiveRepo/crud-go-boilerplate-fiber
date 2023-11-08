package entities

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductUID         string  `gorm:"product_uid" json:"product_uid"`
	ProductName        string  `gorm:"product_name" json:"product_name"`
	PoductPrice        float64 `gorm:"product_price" json:"product_price"`
	ProductStok        float64 `gorm:"product_stok" json:"product_stok"`
	ProductInformation string  `gorm:"product_information" json:"product_information"`
}
