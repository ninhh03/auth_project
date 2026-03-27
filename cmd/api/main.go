package main

import (
	"auth_project/internal/config"
	"auth_project/internal/seed"
	"log"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}
	if err := seed.Seed(db);  err != nil {
		log.Fatalf("seed data failed: %v", err)
	}
}