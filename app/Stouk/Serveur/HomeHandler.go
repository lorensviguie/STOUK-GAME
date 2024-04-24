package serveur

import (
	"data"
	"dice"
	"fmt"
	"html/template"
	"net/http"
	"queue"
	"structure"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	// scoreboard := data.GetScoreBoard()
=======
	scoreboard := data.GetAllScoreBoard()
>>>>>>> 1a67d051f3e6d779aa339ffdd8c3cdda98950ff3
	tmpl, err := template.ParseFiles("./templates/home.html", "./templates/fragments/header.html", "./templates/fragments/footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Println("Scoreboard : ", scoreboard)
	tmpl.Execute(w, nil)
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
	player_Res.Player_data = data.GetAllPlayerDataForGame(playerName)
	var frontData structure.FrontPage
	frontData.Game = player_Res
	frontData.DicePath = dice.BuildDicePathForGame(player_Res)
	tmplFile := "./templates/result.html"
	tmpl, err := template.ParseFiles(tmplFile)
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
