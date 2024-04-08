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

_, err = tx.Exec("INSERT INTO users (Username, Password, Email, Balance, IsAdmin) VALUES (?, ?, ?, ?, ?)", username, hashedPassword, email, 0, 0)
if err != nil {
	return err
}
if err = tx.Commit(); err != nil {
	return err
}
    fmt.Println(username)
    logs.LogToFile("db", "Utilisateur "+username+" ajouté à la base de données avec succès")
    return nil
}

func Login(email, password string) bool {
    db := GetDatabase()
    
    var storedPassword string
    err := db.QueryRow("SELECT Password FROM users WHERE Email = ?", email).Scan(&storedPassword)
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

func GetUsers() ([]structure.Account, error) {
    db := GetDatabase()

    rows, err := db.Query("SELECT ID, Username, Email, Balance FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []structure.Account
    for rows.Next() {
        var user structure.Account
        err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Balance)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func DeleteUser(id string) error {
    db := GetDatabase()
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    _, err = tx.Exec("DELETE FROM account_uuid WHERE ID_USER = ?", id)
    if err != nil {
        return err
    }
    _, err = tx.Exec("DELETE FROM users WHERE ID = ?", id)
    if err != nil {
        return err
    }
    if err = tx.Commit(); err != nil {
        return err
    }
    return nil
}
