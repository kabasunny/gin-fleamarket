package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model // IDを含む構造体

	//ID        uint
	Name        string `gorm:"not null"`
	Price       uint `gorm:"not null"`
	Description string
	SoldOut     bool `gorm:"not null;default:false"` //複数の制約を使用する場合はセミコロンで区切る
}