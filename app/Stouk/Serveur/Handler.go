package serveur

import (
	"data"
	"fmt"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../Stouk/templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tepl, err := template.ParseFiles("../Stouk/templates/loginregister.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tepl.Execute(w, nil)

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if data.Login(username, password) {
			fmt.Println("User " + username + " is logged in")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			fmt.Println("User " + username + " is not logged in")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../Stouk/templates/register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		err := data.Register(username, password)
		if err != nil {
			fmt.Println("User " + username + " is not registered")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			fmt.Println("User " + username + " is registered")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
