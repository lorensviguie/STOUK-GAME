package serveur

import (
	"data"
	"dice"
	"fmt"
	"html/template"
	"logs"
	"net/http"
	"queue"
	"structure"
	"time"
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
	if idUser == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	playerName := data.GetUsernameByUserid(idUser)

	fmt.Println("this user search a Game : ", idUser)

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	logs.LogToFile("serveur", fmt.Sprintf("This user is add to Queue %s,%s", idUser, playerName))
	queue.Add_User_To_Queue(idUser)
	player_Res := queue.CheckTagForUser(idUser)
	logs.LogToFile("serveur", fmt.Sprintf("This user Find a game %s,%s", idUser, playerName))
	player_Res.PlayerName = playerName
	fmt.Println(player_Res)
	player_Res.Player_data = data.GetAllPlayerDataForGame(playerName)
	var frontData structure.FrontPage
	frontData.Game = player_Res
	frontData.DicePath = dice.BuildDicePathForGame(player_Res)
	tmplFile := "./templates/result.html"
	tmpl, err := template.ParseFiles(tmplFile)
	if queue.ContainsID(idUser) {
		for queue.ContainsID(idUser) {
			time.Sleep(1 * time.Second) // Temporisation d'une seconde pour éviter la surcharge du serveur
		}
	}

	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier de modèle HTML :", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, frontData)
	if err != nil {
		fmt.Println("Erreur lors de l'exécution du modèle HTML :", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
