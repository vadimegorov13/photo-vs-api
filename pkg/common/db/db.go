/*
Package db provides functionality for initializing a connection to
the database and performing database migrations.
*/
package db

import (
	"log"

	"github.com/vadimegorov13/photo-vs-api/pkg/common/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDB initializes the database and returns a pointer to a gorm.DB object.
func InitDB() *gorm.DB {
	// Open a connection to the SQLite database located at "posts.db".
	db, err := gorm.Open(sqlite.Open("posts.db"), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	} else {
		// Print a message to the console if the connection was successful.
		log.Println("The database is connected")
	}

	// Automatically migrate the Post model to the database schema.
	db.AutoMigrate(&models.Tournament{})
	log.Println("Dtabase migrated!")

	return db
}
