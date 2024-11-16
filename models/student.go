package models

type Student struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" binding:"required"`
	Age    int    `json:"age" binding:"required"`
	Email  string `json:"email" binding:"required,email"`
	Course string `json:"course" binding:"required"`
}

