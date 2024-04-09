package serveur

import (
	"fmt"
	"net/http"
)

func ServeurInit() {
	fmt.Println("Serveur is running on localhost:8000")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))


	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/dice", DiceHandler)
	http.HandleFunc("/panel-admin", AdminHandler)
	http.HandleFunc("/compte", AccountHandler)
	http.HandleFunc("/play", Playgame)
	
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
