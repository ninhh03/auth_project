package main

import (
	"auth_project/internal/config"
	"auth_project/internal/seed"
)

func main() {
	db := config.ConnectDB()
	seed.Seed(db)
}