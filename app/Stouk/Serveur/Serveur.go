package serveur

import (
	"fmt"
	"net/http"
)

type HtmlData struct {
	IsLoggedIn  bool
	IsModerator bool
	IsAdmin     bool
	PageName    string
	Fragments   map[string]string
}

func ServeurInit() {
	fmt.Println("Serveur is running on localhost:8000")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))


	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/dice", DiceGameWeb)
	
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
