package serveur

import (
	"fmt"
	"net/http"
)

func StoreHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("StoreHandler")
}