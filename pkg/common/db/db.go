/*
Package db provides functionality for initializing a connection to
the database and performing database migrations.
*/
package db

import (
	"fmt"
	"log"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/spf13/viper"
	"github.com/vadimegorov13/photo-vs-api/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the database and returns a pointer to a gorm.DB object.
func InitDB() *gorm.DB {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	projectId := viper.Get("PROJECT_ID").(string)
	projectRegion := viper.Get("PROJECT_REGION").(string)
	dbInstanceName := viper.Get("DB_INSTANCE_NAME").(string)
	dbUser := viper.Get("DB_USER").(string)
	dbName := viper.Get("DB_NAME").(string)
	dbPassword := viper.Get("DB_PASSWORD").(string)

	cloudSQLInstance := fmt.Sprintf("%s:%s:%s", projectId, projectRegion, dbInstanceName)

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", cloudSQLInstance, dbUser, dbName, dbPassword)

	// Open a connection to the SQLite database located at "posts.db".
	db, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "cloudsqlpostgres",
		DSN:        dsn,
	}))

	if err != nil {
		log.Fatalln(err)
	} else {
		// Print a message to the console if the connection was successful.
		log.Println("The database is connected")
	}

	// Automatically migrate the Post model to the database schema.
	if err = db.AutoMigrate(&models.Tournament{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	} else {
		log.Println("Database migrated successfully")
	}

	return db
}
