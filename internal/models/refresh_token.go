package models

import "time"

type RefreshToken struct {
	ID int `json:"id" gorm:"column:id;primaryKey"`
	UserID int `json:"user_id" gorm:"column:user_id;index"`
	Token string `json:"token" gorm:"column:token"`
	ExpiresAt time.Time `json:"expries_at" gorm:"column:expries_at"`
}