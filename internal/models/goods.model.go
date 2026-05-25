// Package models defines persistence models and filters.
package models

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Name  string  `json:"name" gorm:"column:name;type:varchar(100);not null;comment:'Goods name'"`
	Price float64 `json:"price" gorm:"column:price;type:decimal(10,2);not null;comment:'Goods price'"`
	Stock int     `json:"stock" gorm:"column:stock;type:int;not null;comment:'Inventory quantity'"`
}

func (m *Goods) GetID() uint {
	return m.ID
}
func (m *Goods) SetID(id uint) {
	m.ID = id
}
func (m *Goods) TableName() string {
	return "goods"
}
