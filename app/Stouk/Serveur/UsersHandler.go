package serveur

import (
	"data"
	"fmt"
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uuid")
	if err != nil {
		fmt.Println(err)
	}
	IsLoggedIn := false
	if cookie != nil {
		IsLoggedIn = data.CheckAccountUUID(cookie.Value)
		if IsLoggedIn {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	switch r.Method {
	case "GET":
		t, err := template.ParseFiles("./templates/login.html")
		if err != nil {
			fmt.Println(err.Error())
		}
		t.Execute(w, nil)
	case "POST":
		email := r.FormValue("email")
		password := r.FormValue("password")

		if data.Login(email, password) {
			cookie := http.Cookie{
				Name:     "uuid",
				Value:    data.SetAccountUUID(email),
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			fmt.Println("Wrong username or password")
			t, err := template.ParseFiles("./templates/login.html")
			if err != nil {
				fmt.Println(err.Error())
			}
			t.Execute(w, nil)
		}
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uuid")
	if err != nil {
		fmt.Println(err)
	}
	IsLoggedIn := false
	if cookie != nil {
		IsLoggedIn = data.CheckAccountUUID(cookie.Value)
		if IsLoggedIn {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	switch r.Method {
	case "GET":
		t, err := template.ParseFiles("./templates/register.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, nil)
	case "POST":
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		confirmpassword := r.FormValue("confirmpassword")

		if password == confirmpassword {
			data.AddUser(username, password, email)
			cookie := http.Cookie{
				Name:     "uuid",
				Value:    data.SetAccountUUID(email),
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			t, _ := template.ParseFiles("./templates/register.html")

			t.Execute(w, nil)
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("uuid")
	data.RemoveAccountUUID(cookie.Value)
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

