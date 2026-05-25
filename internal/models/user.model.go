package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(30);uniqueIndex;not null"`
	PassWd   string `json:"-" gorm:"not null"`
}

func (m *User) GetID() uint {
	return m.ID
}
func (m *User) SetID(id uint) {
	m.ID = id
}
func (m *User) TableName() string {
	return "users"
}
