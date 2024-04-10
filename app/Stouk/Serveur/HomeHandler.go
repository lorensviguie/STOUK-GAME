package serveur

import (
	"data"
	"fmt"
	"html/template"
	"net/http"
	"queue"
	"structure"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/home.html", "./templates/fragments/header.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := structure.Account{}
	tmpl.Execute(w, data)

}

func SearchGame(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("uuid")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	idUser, _ := data.GetIdByUUID(cookie.Value)
	print("this user search a Game : ", idUser, "\n")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<script>var timer = setTimeout(function() { window.location.href = '/redirect'; }, 5000);</script>")
	queue.Add_User_To_Queue(idUser)
	queue.CheckTagForUser(idUser)
	fmt.Fprintf(w, "<script>clearTimeout(timer); window.location.href = '/results';</script>")
}
