package models

import (

)

type UserRole struct {
	UserID int `json:"user_id" gorm:"column:user_id;primaryKey"`
	RoleID int `json:"role_id" gorm:"column:role_id;primaryKey"`
	
	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASADE"`
	Role Role `gorm:"foreignKey:RoleID;references:ID;constraint:OnDelete:CASADE"`
}