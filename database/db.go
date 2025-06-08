package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=Mm06012000 dbname=mydb port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.ProductInWork{},
	)
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	DB = db
	fmt.Println("Подключение к БД успешно!")
}
