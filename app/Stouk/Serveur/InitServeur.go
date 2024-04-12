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
	http.HandleFunc("/dice", DiceHandler)
	http.HandleFunc("/panel-admin", AdminHandler)
	http.HandleFunc("/compte", AccountHandler)
	http.HandleFunc("/play", Playgame)
	http.HandleFunc("/Queue", SearchGame)
	http.HandleFunc("/update-username", UpdateUsername)
	http.HandleFunc("/update-email", UpdateEmail)
	http.HandleFunc("/update-password", ChangePassword)
	http.HandleFunc("/boutique", StoreHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
