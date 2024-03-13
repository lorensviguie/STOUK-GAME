package data

import (
	"database/sql"
	"fmt"
	"log"
	"logs"

	_ "github.com/go-sql-driver/mysql"
)

const (
	databaseHost     = "localhost"
	databasePort     = "3306"
	databaseUser     = "root"
	databasePassword = ""
	databaseName     = "BA"
)

var db *sql.DB

func InitDatabase() bool {
	var err error
	db, err = sql.Open("mysql", databaseUser+":"+databasePassword+"@tcp("+databaseHost+":"+databasePort+")/"+databaseName+"?parseTime=true")
	if err != nil {
		logs.LogToFile("Error while connecting to the database: "+err.Error(), "error")
	}
	var err1 error
	db, err1 = sql.Open("mysql", databaseUser+":"+databasePassword+"@tcp(ba-db)/"+databaseName+"?parseTime=true")
	if err1 != nil {
		logs.LogToFile("Error while connecting to the database: "+err1.Error(), "error")
		return false
	}
	return true
}

func GetDatabase() *sql.DB {
	return db
}

func CloseDatabase() bool {
	db.Close()
	return true
}

func CreateDB() {
	InitDatabase()
	defer db.Close()

	// Création de la table USERS
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS USERS (
			ID INT AUTO_INCREMENT PRIMARY KEY,
			Username TEXT NOT NULL,
			Password TEXT NOT NULL,
			Balance INT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table USERS created successfully")

	// Création de la table LADDER
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS LADDER (
			ID INT AUTO_INCREMENT PRIMARY KEY,
			ID_USER INT NOT NULL,
			Rank INT NOT NULL,
			MMR INT NOT NULL,
			FOREIGN KEY (ID_USER) REFERENCES USERS(ID)
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table LADDER created successfully")

	// Création de la table RATIO
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS RATIO (
			ID INT AUTO_INCREMENT PRIMARY KEY,
			ID_USER INT NOT NULL,
			Win INT NOT NULL,
			Lose INT NOT NULL,
			RANK_MOYEN INT NOT NULL,
			FOREIGN KEY (ID_USER) REFERENCES USERS(ID)
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table RATIO created successfully")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS DICE (
		ID INT AUTO_INCREMENT PRIMARY KEY,
		Name TEXT NOT NULL,
		Price INT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table DICE created successfully")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS USER_DICE (
		ID INT AUTO_INCREMENT PRIMARY KEY,
		ID_USER INT NOT NULL,
		ID_DICE INT NOT NULL,
		Rank INT NOT NULL,
		FOREIGN KEY (ID_USER) REFERENCES USERS(ID),
		FOREIGN KEY (ID_DICE) REFERENCES DICE(ID)
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table USER_DICE created successfully")

}
