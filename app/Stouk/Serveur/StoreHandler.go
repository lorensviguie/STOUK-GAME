package serveur

import (
	"data"
	"fmt"
	"html/template"
	"net/http"
)

func StoreHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uuid")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if data.CheckAccountUUID(cookie.Value) {
		fmt.Println(cookie.Value)
		user, err := data.GetUserByUUID(cookie.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl, err := template.ParseFiles("./templates/store.html", "./templates/fragments/header.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(user)
		tmpl.Execute(w, user)

	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
