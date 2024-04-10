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

func GetAllUsers() ([]structure.Account, error) {
    db := GetDatabase()

    rows, err := db.Query("SELECT ID, Username, Email, Balance, CreationDate FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []structure.Account
    for rows.Next() {
        var user structure.Account
        err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Balance, &user.CreationDate)
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

func GetUserByUUID(uuid string) (structure.Account, error) {
    db := GetDatabase()
    
    var account structure.Account
    err := db.QueryRow("SELECT USERS.ID, USERS.Username, USERS.Email, USERS.Balance, USERS.CreationDate FROM USERS JOIN ACCOUNT_UUID ON USERS.ID = ACCOUNT_UUID.ID_USER WHERE ACCOUNT_UUID.UUID = ?", uuid).Scan(&account.Id, &account.Username, &account.Email, &account.Balance, &account.CreationDate)
    if err != nil {
        return structure.Account{}, err
    }
    return account, nil
}

func UpdateUsername(uuid, username string) error {
    db := GetDatabase()
    
    _, err := db.Exec("UPDATE USERS JOIN ACCOUNT_UUID ON USERS.ID = ACCOUNT_UUID.ID_USER SET USERS.Username = ? WHERE ACCOUNT_UUID.UUID = ?", username, uuid)
    if err != nil {
        return err
    }
    return nil
}

func UpdateEmail(uuid, email string) error {
    db := GetDatabase()
    
    _, err := db.Exec("UPDATE USERS JOIN ACCOUNT_UUID ON USERS.ID = ACCOUNT_UUID.ID_USER SET USERS.Email = ? WHERE ACCOUNT_UUID.UUID = ?", email, uuid)
    if err != nil {
        return err
    }
    return nil
}

func ChangePassword(uuid, password string) error {
    db := GetDatabase()
    
    hashedPassword, err := HashPassword(password)
    if err != nil {
        return err
    }
    
    _, err = db.Exec("UPDATE USERS JOIN ACCOUNT_UUID ON USERS.ID = ACCOUNT_UUID.ID_USER SET USERS.Password = ? WHERE ACCOUNT_UUID.UUID = ?", hashedPassword, uuid)
    if err != nil {
        return err
    }
    return nil
}

func CheckPasswordByUUID(uuid, password string) bool {
    db := GetDatabase()
    
    var storedPassword string
    err := db.QueryRow("SELECT USERS.Password FROM USERS JOIN ACCOUNT_UUID ON USERS.ID = ACCOUNT_UUID.ID_USER WHERE ACCOUNT_UUID.UUID = ?", uuid).Scan(&storedPassword)
    if err != nil {
        return false
    }
    err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
    if err != nil {
        return false
    }
    return true
    
    
}
