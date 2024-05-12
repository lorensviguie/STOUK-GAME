package serveur

import (
	"data"
	"fmt"
	"html/template"
	"net/http"
	"structure"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uuid")
	if err != nil {
		HomeHandler(w, r)
		return
	}
	if data.IsAdmin(cookie.Value) {
		users, err := data.GetAllUsers()
		fmt.Println(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl, err := template.ParseFiles("./templates/admin.html", "./templates/fragments/header.html", "./templates/fragments/footer.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if r.Method == "POST" {
			r.ParseForm()
			id := r.FormValue("id")
			action := r.FormValue("action")
			balance := r.FormValue("balance")
			fmt.Println("ID:", id, "Action:", action, "Balance:", balance)
		
			switch action {
			case "delete":
				data.DeleteUser(id)
			case "add_balance":
				if balance != "" {
					data.UpdateBalance(id, balance)
				}
			case "remove_admin":
				data.DeleteAdmin(id, "")
			case "set_admin":
				data.SetAdmin(id, "")
			}
		
			http.Redirect(w, r, "/panel-admin", http.StatusSeeOther)
		}
		

		tmpl.Execute(w, structure.AdminData{Users: users})
	} else {
		HomeHandler(w, r)
		return
	}
}
