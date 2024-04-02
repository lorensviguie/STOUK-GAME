package serveur

import (
	"fmt"
	"net/http"
)

func DiceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DiceHandler")
}