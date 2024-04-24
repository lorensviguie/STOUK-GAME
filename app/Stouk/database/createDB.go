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
	DB, err = sql.Open("mysql", "root:STOUK@tcp(STOUK-GAME:3306)/Stouk-GAME")
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
			Email TEXT NOT NULL,
			Balance INT NOT NULL,
			IsAdmin BOOLEAN NOT NULL,
			CreationDate DATE NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table ACCOUNT created successfully")

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
		Price INT NOT NULL,
		Path TEXT NOT NULL
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

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS ACCOUNT_UUID (
		ID INT AUTO_INCREMENT PRIMARY KEY,
		ID_USER INT NOT NULL,
		UUID TEXT NOT NULL,
		CREATE_DATE DATE NOT NULL,
		FOREIGN KEY (ID_USER) REFERENCES USERS(ID)
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table ACCOUNT_UUID created successfully")

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS PROFIL_PICTURE (
		ID INT AUTO_INCREMENT PRIMARY KEY,
		ID_USER INT NOT NULL,
		PICTURE TEXT NOT NULL,
		FOREIGN KEY (ID_USER) REFERENCES USERS(ID)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table PROFIL_PICTURE created successfully")

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS PRICE (
		ID INT AUTO_INCREMENT PRIMARY KEY,
		PRICE INT NOT NULL
	);
	`)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table PRICE created successfully")


	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS HISTORIQUE (
	ID INT AUTO_INCREMENT PRIMARY KEY,
	ID_PLAYER_1 INT NOT NULL,
	ID_PLAYER_2 INT NOT NULL,
	ID_PLAYER_WIN INT NOT NULL,
	PLAYER_1_WIN INT NOT NULL,
	PLAYER_2_WIN INT NOT NULL,
	PLAYER_1_LP_MOD INT NOT NULL,
	PLAYER_2_LP_MOD INT NOT NULL,
	FOREIGN KEY (ID_PLAYER_1) REFERENCES USERS(ID),
	FOREIGN KEY (ID_PLAYER_2) REFERENCES USERS(ID)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table HISTORIQUE created successfully")

	Add_Dice("BaseDice", 10, "./static/images/de1.svg")
	Add_Dice("NormalDice", 10, "./static/images/de2.svg")
	Add_Dice("ParaboleDice", 10, "./static/images/de3.svg")
	Add_Dice("PowerDice", 10, "./static/images/de4.svg")
	Add_Dice("ScaleDice", 10, "./static/images/de5.svg")
	Add_Dice("UnscaleDice", 10, "./static/images/de6.svg")
	Add_Dice("RankDice", 10, "./static/images/de7.svg")

	Add_Price(100)
	Add_Price(200)
	Add_Price(320)
	Add_Price(450)
	Add_Price(650)
	Add_Price(1000)
	Add_Price(1800)
	Add_Price(3700)
	Add_Price(6000)
}
