package config

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GalaSetup() (*gorm.DB, error) {
	// GalaSetup : Set Connection to MySQL !
	strict := "root:secretapps@tcp(sample_golang_apps_compose)/sample_golang_apps?charset=utf8mb4&parseTime=True"

	database, err := gorm.Open(mysql.Open(strict), &gorm.Config{})
	if err != nil {
		log.Fatalf("Fail to Connect Database: %v", err)
		return nil, err
	}

	sql_database, err := database.DB()

	if err != nil {
		log.Fatalf("Fail to Retrieve DB : %v", err)
	}

	sql_database.SetMaxIdleConns(10)
	sql_database.SetMaxOpenConns(100)
	sql_database.SetConnMaxLifetime(30 * time.Minute)

	return database, nil
}

// docker run --name some-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:tag
