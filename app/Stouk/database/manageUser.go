package data

import (
	"database/sql"
	"fmt"
	"logs"
	"structure"

	"golang.org/x/crypto/bcrypt"
)

func AddUser(username, password, email string) error {
	db := GetDatabase()
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error starting transaction:", err)
		return err
	}
	defer tx.Rollback() // Rollback in case of any errors

	hashedPassword, err := HashPassword(password) // Assuming HashPassword hashes the password
	if err != nil {
		return err
	}

	result, err := tx.Exec("INSERT INTO USERS (Username, Password, Email, Balance) VALUES (?, ?, ?, ?)", username, hashedPassword, email, 0)
	if err != nil {
		fmt.Println("Error inserting user:", err)
		return err
	}

	// Obtenir l'ID de l'utilisateur inséré
	userID, _ := result.LastInsertId()

	// Insert user entry into LADDER table with default values
	_, err = tx.Exec("INSERT INTO LADDER (ID_USER, Rank, MMR) VALUES (?, ?, ?)", userID, 1000, 1000) // Adjust default Rank and MMR as needed
	if err != nil {
		fmt.Println("Error inserting user into LADDER:", err)
		return err
	}

	// Insert user entry into RATIO table with default values
	_, err = tx.Exec("INSERT INTO RATIO (ID_USER, Win, Lose, RANK_MOYEN) VALUES (?, ?, ?, ?)", userID, 0, 0, 1000) // Adjust default Win, Lose, and RANK_MOYEN as needed
	if err != nil {
		fmt.Println("Error inserting user into RATIO:", err)
		return err
	}

	if err = tx.Commit(); err != nil {
		fmt.Println("Error committing transaction:", err)
		return err
	}

	fmt.Println(username)
	logs.LogToFile("db", "Utilisateur "+username+" ajouté à la base de données avec succès")
	return nil
}

func Login(username, password string) bool {
	db := GetDatabase()

	var storedPassword string
	err := db.QueryRow("SELECT Password FROM users WHERE Username = ?", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false
		}
		return false
	}
	return true
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GetAccountByEmail(email string, withDefer bool) structure.Account {
	db := GetDatabase()

	var account structure.Account
	err := db.QueryRow("SELECT ID, Username, Email FROM users WHERE Email = ?", email).Scan(&account.Id, &account.Username, &account.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return structure.Account{}
		}
		return structure.Account{}
	}
	return account
}

func GetUserIDByUsername(username string) (int64, error) {
    db := GetDatabase() // Supposons que vous ayez une fonction GetDatabase() qui retourne une connexion à la base de données

    var userID int64
    err := db.QueryRow("SELECT ID FROM USERS WHERE Username = ?", username).Scan(&userID)
    if err != nil {
        if err == sql.ErrNoRows {
            return 0, fmt.Errorf("utilisateur avec le nom d'utilisateur %s non trouvé", username)
        }
        return 0, err
    }

    return userID, nil
}