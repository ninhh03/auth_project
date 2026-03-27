package seed

import (
	"auth_project/internal/models"
	"auth_project/pkg/utils"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	roles := []models.Role{
		{Name: "admin"},
		{Name: "user"},
	}
	for _, r := range roles {
		db.FirstOrCreate(&r, models.Role{Name: r.Name})
	}

	hash1, _ := utils.HashPassword("123")
	hash2, _ := utils.HashPassword("456")
	hash3, _ := utils.HashPassword("789")
	hash4, _ := utils.HashPassword("101112")
	users := []models.User{
		{
			FullName: "Nguyễn Dương Ninh",
			Age:      23,
			Address:  "Hà Nội",
			Email:    "ndninh.040303@gmai.com",
			Password: hash1,
		},
		{
			FullName: "Nguyễn Hữu Huy",
			Age:      33,
			Address:  "Hà Nội",
			Email:    "nhhuy@gmai.com",
			Password: hash2,
		},
		{
			FullName: "Ngô Lan Phương",
			Age:      20,
			Address:  "Hà Nội",
			Email:    "nlphuong@gmai.com",
			Password: hash3,
		},
		{
			FullName: "Nguyễn Công Minh Nghĩa",
			Age:      23,
			Address:  "Hà Nội",
			Email:    "enciemen@gmai.com",
			Password: hash4,
		},
		{
			FullName: "Nguyễn Công Minh Nghĩa",
			Age:      23,
			Address:  "Hà Nội",
			Email:    "enciemen@gmai.com",
			Password: hash4,
		},
	}
	for _, u := range users {
		db.FirstOrCreate(&u, models.User{Email: u.Email})
	}

	var user1 models.User
	db.Where("email = ?", "ndninh.040303@gmai.com").First(&user1)
	var role1 models.Role
	db.Where("name = ?", "admin").First(&role1)
	var role2 models.Role
	db.Where("name = ?", "user").First(&role2)
	user_roles := []models.UserRole {
		{
			UserID: user1.ID,
			RoleID: role1.ID,
		},
		{
			UserID: user1.ID,
			RoleID: role2.ID,
		},
	}
	for _,ur := range user_roles {
		db.FirstOrCreate(&ur, models.UserRole{
			UserID: ur.UserID,
			RoleID: ur.RoleID,
		})
	}
}
