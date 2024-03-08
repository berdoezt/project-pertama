package lib

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	connectionString := "host=localhost port=5432 user=postgres password=password dbname=hacktiv sslmode=disable"
	return gorm.Open(postgres.Open(connectionString), &gorm.Config{})
}
