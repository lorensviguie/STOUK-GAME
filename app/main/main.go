package main

import (
	"data"
	"serveur"
)

func main() {
	serveur.ServeurInit()
	data.InitDatabase()
}

