package main

import (
	"github.com/azicussdu/go-lms/cmd/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	// Database connection details
	dsn := "host=localhost user=my_user password=my_pass dbname=golms port=5438 sslmode=disable"

	// Open a connection to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect to the database")
	}

	// Auto migrate the schema
	if err = db.AutoMigrate(&models.User{}, &models.Category{}, &models.Course{}, &models.StudentCourse{}, &models.Author{}, &models.Topic{}, &models.Lesson{}); err != nil {
		log.Println(err.Error())
		return
	}

	//user := models.User{Email: "azicus.sdu@gmail.com", PasswordHash: "sadwr34rseasfr43sgw4s"}
	//result := db.Create(&user)
	//
	//if result.Error != nil {
	//	log.Println(result.Error.Error())
	//	return
	//}

	log.Println("Reached IT ----------- ")
}
