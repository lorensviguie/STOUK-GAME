package serveur

import (
	"data"
	"html/template"
	"net/http"
	"strconv"
	"structure"
)

func StoreHandler(w http.ResponseWriter, r *http.Request) {
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
		dices, err := data.GetUserDice(user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl, err := template.ParseFiles("./templates/store.html", "./templates/fragments/header.html", "./templates/fragments/footer.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, structure.StoreData{User: []structure.Account{user}, Dices: dices})

	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func RankUp(w http.ResponseWriter, r *http.Request) {
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

		diceID := r.FormValue("diceNumber")
		dice, err := data.GetUserDice(user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var selectedDice structure.Dice
		for _, d := range dice {
			diceIDInt, err := strconv.Atoi(diceID)
			if err != nil {
				return
			}
			if d.Dice == diceIDInt {
				selectedDice = d
				break
			}
		}

		if user.Balance >= selectedDice.Price {
			err := data.UpdateRank(user.Id, diceID, selectedDice.Price)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		http.Redirect(w, r, "/boutique", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/boutique", http.StatusSeeOther)
		return
	}
}
