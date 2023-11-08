package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderUID  string `gorm:"order_uid" json:"order_uid"`
	ProductID uint64 `gorm:"product_id" json:"product_id"`
}
