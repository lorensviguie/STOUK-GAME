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
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/panel-admin", AdminHandler)
	http.HandleFunc("/compte", AccountHandler)
	http.HandleFunc("/queue", SearchGame)
	http.HandleFunc("/boutique", StoreHandler)
	http.HandleFunc("/update-compte", UpdateAccount)
	http.HandleFunc("/update-profil-picture", UpdateProfilPicture)
	http.HandleFunc("/rankup", RankUp)
	
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
