package main

import (
	"data"
)

func main() {
	data.CreateDB()
	data.Add_Dice("base",10)
	data.CreateUser("farkas", "test")
	data.CreateUser("coin", "test")
}
