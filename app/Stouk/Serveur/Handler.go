package serveur

import (
	"data"
	"fmt"
	"html/template"
	"net/http"
	"structure"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../Stouk/templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := structure.Account{}
	tmpl.Execute(w, data)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userInfo structure.Account
	if r.Method == "POST" {
		fmt.Println("Logging in user")
		userInfo.Username = r.FormValue("username")
		userInfo.Password = r.FormValue("password")
		fmt.Println(userInfo.Username, userInfo.Password)
		err := data.Login(userInfo.Username, userInfo.Password)
		if !err {
			fmt.Println("User " + userInfo.Username + " is logged in")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			fmt.Println("User " + userInfo.Username + " is not logged in")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
	fmt.Println("Page loaded")
	temp, err := template.ParseFiles("../Stouk/templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var userInfo structure.Account
	if r.Method == "POST" {
		fmt.Println("Registering user")
		userInfo.Username = r.FormValue("username")
		userInfo.Email = r.FormValue("email")
		userInfo.Password = r.FormValue("password")
		fmt.Println(userInfo.Username, userInfo.Password)
		registerErr := data.Register(userInfo.Username, userInfo.Password, userInfo.Email)
		uuidErr := data.SetAccountUUID(userInfo.Email)
		if registerErr == nil && uuidErr {
			fmt.Println("User " + userInfo.Username + " is registered")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			fmt.Println("User " + userInfo.Username + " is not registered")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
	fmt.Println("Page loaded")
	temp, err := template.ParseFiles("../Stouk/templates/register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}
