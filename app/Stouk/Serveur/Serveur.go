package serveur

import (
	"fmt"
	"net/http"
)

type HtmlData struct {
	IsLoggedIn  bool
	IsModerator bool
	IsAdmin     bool
	PageName string
	Fragments map[string]string
}

func ServeurInit() { 
	fmt.Println("Serveur is running on localhost:8000")

	http.HandleFunc("/", Homehandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.ListenAndServe(":8000", nil)

}
