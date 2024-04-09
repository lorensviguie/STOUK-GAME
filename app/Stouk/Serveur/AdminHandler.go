package serveur

import (
	"data"
	"html/template"
	"net/http"
	"structure"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("uuid")
    if err != nil {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    if data.IsAdmin(cookie.Value) {
        users, err := data.GetUsers()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        tmpl, err := template.ParseFiles("./templates/admin.html", "./templates/fragments/header.html")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Handle POST request separately to avoid interference with template execution
        if r.Method == "POST" {
            r.ParseForm()
            id := r.FormValue("id")
            data.DeleteUser(id)
			http.Redirect(w, r, "/panel-admin", http.StatusSeeOther)
            return
        }

        tmpl.Execute(w, structure.AdminData{Users: users})
    } else {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
}
