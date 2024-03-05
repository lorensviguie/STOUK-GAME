package database

import (
	logs "STOUK/Stouk/logs"
	"fmt"
)

func CreateUser(username string, password string) {
	db := ConnectToDb()
	_, err := db.Exec("INSERT INTO users (Username, Password, Balance) VALUES (?, ?, ?)", username, password, 0)
	if err != nil {
		logs.LogToFile("db", err.Error())
		panic(err)
	}
	fmt.Println(username)
	logs.LogToFile("db", "Users "+username+" add to db with succes")
}
