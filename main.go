package main

import (
	"go-back/entity"
	"go-back/router"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// dbを作成します

	db := dbInit()

	// dbをmigrateします
	db.AutoMigrate(&entity.Article{})
	router.GetRouter(db)

}

func dbInit() *gorm.DB {

	dsn := "host=localhost user=bloguser password=bloguser dbname=blog port=15432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
