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
	_, err := db.Exec("INSERT INTO account_uuid (id_user, uuid, create_date) VALUES (?, ?, ?)", accountId, uuid, now)
	if err != nil {
		fmt.Println(err)
	}
	return uuid
}

func CheckAccountUUID(uuid string) bool {
	db := GetDatabase()
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM account_uuid WHERE uuid = ?", uuid).Scan(&count)
	if err != nil {
		fmt.Println(err)
	}
	return count == 1
}

func RemoveAccountUUID(uuid string) {
	db := GetDatabase()
	_, err := db.Exec("DELETE FROM account_uuid WHERE uuid = ?", uuid)
	if err != nil {
		fmt.Println(err)
	}
}