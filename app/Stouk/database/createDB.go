package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	databaseUser     = "root"
// 	databasePassword = "STOUK"
// 	databaseName     = "Stouk-GAME"
// 	databaseHost     = "localhost"
// 	databasePort     = "3306"
// )

var DB *sql.DB

func InitDatabase() bool {
	var err error
	DB, err = sql.Open("mysql", "root:STOUK@tcp(localhost:3306)/Stouk-GAME")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	return true
}

func GetDatabase() *sql.DB {
	return DB
}

func CloseDatabase() bool {
	DB.Close()
	return true
}

func CreateDB() {
	InitDatabase()
	// Création de la table USERS
	_, err := DB.Exec(`
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
	_, err = DB.Exec(`
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
	_, err = DB.Exec(`
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

	_, err = DB.Exec(`
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

	_, err = DB.Exec(`
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
