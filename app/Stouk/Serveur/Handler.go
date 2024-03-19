package serveur 

import (
	"net/http"
	"html/template"
	"fmt"
	"data"
)

func Homehandler(w http.ResponseWriter, r *http.Request) {
	// the path of the page is /app/Stouk/src/home.html
	tmpl, err := template.ParseFiles("../Stouk/src/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tepl, err := template.ParseFiles("../Stouk/src/loginregister.html")
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
	tmpl, err := template.ParseFiles("../Stouk/src/register.html")
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