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
	tmpl, err := template.ParseFiles("../Stouk/templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := structure.Account{}
	tmpl.Execute(w, data)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var userInfo structure.Account
	if r.Method == "POST" {
		fmt.Println("Logging in user")
		userInfo.Username = r.FormValue("username")
		userInfo.Password = r.FormValue("password")
		fmt.Println(userInfo.Username, userInfo.Password)
		err := data.Login(userInfo.Username, userInfo.Password)
		if !err {
			fmt.Println("User " + userInfo.Username + " is logged in")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			fmt.Println("User " + userInfo.Username + " is not logged in")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
	fmt.Println("Page loaded")
	temp, err := template.ParseFiles("../Stouk/templates/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		t, err := template.ParseFiles("../Stouk/templates/register.html")
		if err != nil {
			fmt.Println(err)
		}

		var IsLoggedIn bool
		cookie, err := r.Cookie("uuid")
		if err != nil {
			fmt.Println(err)
			IsLoggedIn = false
		} else {
			IsLoggedIn = data.CheckAccountUUID(cookie.Value)
			if IsLoggedIn {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}

		// prepareDataWithFragments(&data)
		t.Execute(w, nil)
	case "POST":
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		confirmpassword := r.FormValue("confirmpassword")

		fmt.Println(password)
		fmt.Println(confirmpassword)

		bool := confirmpassword == password
		fmt.Println(bool)

		if password == confirmpassword {
			data.AddUser(username, password, email)

			cookie := http.Cookie{
				Name:  "uuid",
				Value: data.SetAccountUUID(email),
			}
			http.SetCookie(w, &cookie)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			fmt.Println("Passwords do not match HERE")
			t, err := template.ParseFiles("../Stouk/templates/register.html")
			if err != nil {
				fmt.Println(err)
			}
			// prepareDataWithFragments(&data)
			t.Execute(w, nil)
		}
	}
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
