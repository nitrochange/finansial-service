package dao

import (
	"finansial-service/main/models"
	"finansial-service/main/models/requiests"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	_ "github.com/lib/pq"
	"log"
	"os"
)

//DB
var DBConnect *pg.DB

// Connecting to db
func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "admin",
		Password: "password",
		Addr:     "localhost:5432",
		Database: "postgres",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")

	//USERS TABLE CREATION
	opts1 := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&models.User{}, opts1)
	if createError != nil {
		log.Printf("Error while creating USERS table, Reason: %v\n", createError)
	}
	log.Printf("Users table created")

	//TRANSACTIONS TABLE CREATION
	opts2 := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError = db.CreateTable(&requiests.Transaction{}, opts2)
	if createError != nil {
		log.Printf("Error while creating TRANSACTIONS table, Reason: %v\n", createError)
	}
	log.Printf("TRANSACTIONS table created")

	DBConnect = db
	return db
}
