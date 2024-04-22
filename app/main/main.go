package main

import (
	"data"
	// "queue"
	"serveur"
)

func main() {
	data.CreateDB()
	// go queue.ManageQueue()
	serveur.ServeurInit()

}
