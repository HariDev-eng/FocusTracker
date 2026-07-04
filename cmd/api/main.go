package main

import (
	"focustracker/internal/handlers"
	"focustracker/internal/models"
	"focustracker/internal/router"
	"log"

	"focustracker/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&models.User{}, &models.Task{}, &models.Completion{}); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	seedDefaultUser(db)
	log.Println("migrations complete")

	log.Println("connected to database successfully")
	_ = db

	h := handlers.NewHandler(db)
	r := router.New(h)
	log.Printf("server starting on port %s", cfg.Port)
	err = r.Run(":" + cfg.Port)
	if err != nil {
		return
	}
}

func seedDefaultUser(db *gorm.DB) {
	var count int64
	db.Model(&models.User{}).Count(&count)
	if count == 0 {
		db.Create(&models.User{Email: "hp@local"})
	}
}
