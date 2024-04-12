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
	playerName := data.GetUsernameByUserid(idUser)

	fmt.Println("this user search a Game : ", idUser)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	queue.Add_User_To_Queue(idUser)
	player_Res := queue.CheckTagForUser(idUser)
	player_Res.PlayerName = playerName
	tmplFile := "./templates/result.html"
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier de modèle HTML :", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, player_Res)
	if err != nil {
		fmt.Println("Erreur lors de l'exécution du modèle HTML :", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
