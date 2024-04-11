package serveur

import (
	"data"
	"fmt"
	"html/template"
	"net/http"

)

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uuid")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if data.CheckAccountUUID(cookie.Value) {
		user, err := data.GetUserByUUID(cookie.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(user)

		tmpl, err := template.ParseFiles("./templates/account.html", "./templates/fragments/header.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, user)
		
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func UpdateUsername(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uuid")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if data.CheckAccountUUID(cookie.Value) {
		username := r.FormValue("username")
		data.UpdateUsername(cookie.Value, username)
		http.Redirect(w, r, "/compte", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func UpdateEmail(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uuid")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if data.CheckAccountUUID(cookie.Value) {
		email := r.FormValue("email")
		data.UpdateEmail(cookie.Value, email)
		http.Redirect(w, r, "/compte", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("uuid")
    if err != nil {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    if data.CheckAccountUUID(cookie.Value) {
        oldPassword := r.FormValue("password")
        
        if !data.CheckPasswordByUUID(cookie.Value, oldPassword) {
            http.Redirect(w, r, "/compte", http.StatusSeeOther)
            return
        } else {
        newPassword := r.FormValue("newpassword")
        data.ChangePassword(cookie.Value, newPassword)
        http.Redirect(w, r, "/compte", http.StatusSeeOther)
        }
    } else {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
}
