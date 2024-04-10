package serveur

import (
	"data"
	"html/template"
	"net/http"
	"structure"
)

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uuid")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if data.CheckAccountUUID(cookie.Value) {
		tmpl, err := template.ParseFiles("./templates/account.html", "./templates/fragments/header.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := structure.Account{}
		tmpl.Execute(w, data)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
}
