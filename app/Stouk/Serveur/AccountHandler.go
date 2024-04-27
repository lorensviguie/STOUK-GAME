package serveur

import (
	"data"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

func AccountHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uuid")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if data.CheckAccountUUID(cookie.Value) {
		user, err := data.GetUserByUUID(cookie.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		Compte := data.GetALLDataForHistorique(user)
		tmpl, err := template.ParseFiles("./templates/account.html", "./templates/fragments/header.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(Compte)
		tmpl.Execute(w, Compte)

	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uuid")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if data.CheckAccountUUID(cookie.Value) {
		username := r.FormValue("username")
		if username == "" {
			http.Redirect(w, r, "/compte", http.StatusSeeOther)
			return
		}
		data.UpdateUsername(cookie.Value, username)
		http.Redirect(w, r, "/compte", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if data.CheckAccountUUID(cookie.Value) {
		email := r.FormValue("email")
		if email == "" {
			http.Redirect(w, r, "/compte", http.StatusSeeOther)
			return
		}
		data.UpdateEmail(cookie.Value, email)
		http.Redirect(w, r, "/compte", http.StatusSeeOther)
	} else {
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
			if newPassword == "" {
				http.Redirect(w, r, "/compte", http.StatusSeeOther)
				return
			}
			data.ChangePassword(cookie.Value, newPassword)
			http.Redirect(w, r, "/compte", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func UpdateProfilPicture(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uuid")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if data.CheckAccountUUID(cookie.Value) {
		r.ParseMultipartForm(10 << 20)
		fmt.Println("Oui")
		file, handler, err := r.FormFile("profil-picture")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		filePath := "./static/images/profilpicture/" + strings.Split(handler.Filename, ".")[0] + "_" + cookie.Value + ".png"
		f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Error Saving the File")
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		err = data.UpdateProfilPicture(cookie.Value, filePath)
		if err != nil {
			fmt.Println("Error Updating Profile Picture in Database")
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, "/compte", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
