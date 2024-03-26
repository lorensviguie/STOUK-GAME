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
	return err
}
defer tx.Rollback()

hashedPassword, err := HashPassword(password)
if err != nil {
	return err
}

_, err = tx.Exec("INSERT INTO users (Username, Password, Email, Balance) VALUES (?, ?, ?, ?)", username, hashedPassword, email, 0)
if err != nil {
	return err
}
if err = tx.Commit(); err != nil {
	return err
}
// set new account uuid



    
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

