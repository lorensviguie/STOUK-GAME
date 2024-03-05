package database

import (
	"fmt"
	"log"
)

func CreateDB() {
	var db = ConnectToDb()
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
}
