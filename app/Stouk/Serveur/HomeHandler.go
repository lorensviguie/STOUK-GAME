package serveur 

import (
	"data"
	"net/http"
	"html/template"
)

func Homehandler(w http.ResponseWriter, r *http.Request) {
	var account data.Account

	htmlData := HtmlData{
		PageName: "Home",
		IsLoggedIn: false,
		IsModerator: false,
		IsAdmin: false,
	}

	cookie, err := r.Cookie("uuid")
	if err != nil {
		htmlData.IsLoggedIn = false
	} else {
		IsLoggedIn := data.CheckAccountUUID(cookie.Value)
		htmlData.IsLoggedIn = IsLoggedIn
		account, err = data.GetAccount(data.GetAccountIdByUUID(cookie.Value))
		if err != nil {
		}
		htmlData.IsModerator = account.IsModerator
		htmlData.IsAdmin = account.IsAdmin
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
	}

	prepareDataWithFragments(&htmlData)
	t.Execute(w, htmlData)
}