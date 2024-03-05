package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DataBase *sql.DB

func ConnectToDb() *sql.DB {
	db, err := sql.Open("mysql", "root:STOUK@tcp(stouk-db:3306)/Stouk-GAME")
	if err != nil {
		log.Fatal(err)
	}

	// Vérification de la connexion à la base de données
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	return db
}
