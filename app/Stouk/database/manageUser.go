package database

import (
	logs "STOUK/Stouk/logs"
	"fmt"
)

func CreateUser(username string, key string) {
	db := ConnectToDb()
	_, err := db.Exec("INSERT INTO users (username, key) VALUES (?, ?)", username, key)
	if err != nil {
		logs.LogToFile("db", err.Error())
		panic(err)
	}
	fmt.Println(username)
	logs.LogToFile("db", "Users "+username+" add to db with succes")
}
