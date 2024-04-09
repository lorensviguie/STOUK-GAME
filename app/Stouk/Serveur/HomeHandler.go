package serveur

import (
	"data"
	"dice"
	"fmt"
	"html/template"
	"ladder"
	"net/http"
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

}

func DiceGameWeb(w http.ResponseWriter, r *http.Request) {
	print("\n\n--------------\nCOIN COIN COIN COIN\n----------------\n\n")
	player1 := data.GetAllPlayerDataForGame("toto")
	player2 := data.GetAllPlayerDataForGame("lolo")
	Player1res, Player2res := dice.Dice_Game(player1, player2)
	fmt.Println(Player1res)
	fmt.Println(Player2res)
	if Player1res.Game_res.Player1_Win > Player1res.Game_res.Player2_Win {
		Player1res.Player_data.Win += 1
		Player2res.Player_data.Lose += 1
		Player1res.Player_data.MMR += 20
		Player2res.Player_data.MMR -= 20
		Player1res.Player_data.RankMoyen = ((Player1res.Player_data.Win+Player1res.Player_data.Lose)*Player1res.Player_data.RankMoyen + Player2res.Player_data.Rank) / (Player1res.Player_data.Win + Player1res.Player_data.Lose + 1)
		Player2res.Player_data.RankMoyen = ((Player2res.Player_data.Win+Player2res.Player_data.Lose)*Player2res.Player_data.RankMoyen + Player1res.Player_data.Rank) / (Player2res.Player_data.Win + Player2res.Player_data.Lose + 1)
		ladder.UpdateRankforPlayer(Player1res.Player_data, true)
		ladder.UpdateRankforPlayer(Player2res.Player_data, false)
	} else {
		Player2res.Player_data.Win += 1
		Player1res.Player_data.Lose += 1
		Player1res.Player_data.MMR -= 20
		Player2res.Player_data.MMR += 20
		Player1res.Player_data.RankMoyen = ((Player1res.Player_data.Win+Player1res.Player_data.Lose)*Player1res.Player_data.RankMoyen + Player2res.Player_data.Rank) / (Player1res.Player_data.Win + Player1res.Player_data.Lose + 1)
		Player2res.Player_data.RankMoyen = ((Player2res.Player_data.Win+Player2res.Player_data.Lose)*Player2res.Player_data.RankMoyen + Player1res.Player_data.Rank) / (Player2res.Player_data.Win + Player2res.Player_data.Lose + 1)
		ladder.UpdateRankforPlayer(Player1res.Player_data, false)
		ladder.UpdateRankforPlayer(Player2res.Player_data, true)
	}
	player1 = data.GetAllPlayerDataForGame("toto")
	player2 = data.GetAllPlayerDataForGame("lolo")
	fmt.Println("Player toto Rank : ", player1.Rank, " RankMoyen ", player1.RankMoyen, " Avec ", player1.Win, " Victoire et ", player1.Lose, "Defaite ceux qui lui fait un mmr de ", player1.MMR)
	fmt.Println("Player lolo Rank : ", player2.Rank, " RankMoyen ", player2.RankMoyen, " Avec ", player2.Win, " Victoire et ", player2.Lose, "Defaite ceux qui lui fait un mmr de ", player2.MMR)
}
