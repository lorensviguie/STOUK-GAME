package data

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofrs/uuid"
)

func generateUUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Println(err)
	}
	return uuid.String()
}

func SetAccountUUID(email string) string {
	db := GetDatabase()
	accountId := GetAccountByEmail(email, false).Id

	now := time.Now().UTC()

	uuid := generateUUID()
	_, err := db.Exec("INSERT INTO ACCOUNT_UUID (id_user, uuid, create_date) VALUES (?, ?, ?)", accountId, uuid, now)
	if err != nil {
		fmt.Println(err)
	}
	return uuid
}

func CheckAccountUUID(uuid string) bool {
	db := GetDatabase()
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM ACCOUNT_UUID WHERE uuid = ?", uuid).Scan(&count)
	if err != nil {
		fmt.Println(err)
	}
	return count == 1
}

func RemoveAccountUUID(uuid string) {
	db := GetDatabase()
	_, err := db.Exec("DELETE FROM ACCOUNT_UUID WHERE uuid = ?", uuid)
	if err != nil {
		fmt.Println(err)
	}
}

func IsAdmin(uuid string) bool {
	db := GetDatabase()
	var isadmin int
	err := db.QueryRow("SELECT USERS.IsAdmin FROM USERS JOIN ACCOUNT_UUID ON USERS.ID = ACCOUNT_UUID.ID_USER WHERE ACCOUNT_UUID.UUID = ?", uuid).Scan(&isadmin)
	if err != nil {
		fmt.Println(err)
	}
	return isadmin == 1
}

func GetIdByUUID(uuid string) (int, error) {
    db := GetDatabase()

    var id int
    err := db.QueryRow("SELECT ID_USER FROM ACCOUNT_UUID WHERE UUID = ?", uuid).Scan(&id)
    if err != nil {
        return 0, err
    }
    return id, nil
}