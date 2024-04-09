package main

import (
	"data"
	"serveur"
)

func main() {
	data.CreateDB()
	serveur.ServeurInit()
}
