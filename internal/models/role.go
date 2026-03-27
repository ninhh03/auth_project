package models

import (

)

type Role struct {
	ID int `json:"id" gorm:"column:id;primaryKey"`
	Name string `json:"name" gorm:"column:name"`
}