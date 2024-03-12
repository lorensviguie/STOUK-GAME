package main

import (
	db "STOUK/Stouk/database"
	"serveur"
)

func main() {
	serveur.ServeurInit()
	db.CreateDB()
}

