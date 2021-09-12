package model

type User struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `json:"name" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Age      int    `json:"age" gorm:"not null"`
}