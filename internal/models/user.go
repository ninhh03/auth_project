package models

import (

)

type User struct {
	ID int `json:"id" gorm:"column:id;primaryKey"`
	FullName string `json:"fullname" gorm:"column:fullname"`
	Age int `json:"age" gorm:"column:age"`
	Address string `json:"address" gorm:"column:address"`
	Email string `json:"email" gorm:"column:email;uniqueIndex"`
	Password string `json:"-" gorm:"column:password"`
}