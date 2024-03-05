package serveur

import (
	"fmt"
	"net/http"
)

func ServeurInit() { 
	fmt.Println("Serveur is running on localhost:8000")
	http.HandleFunc("/", Homehandler)
	http.ListenAndServe(":8000", nil)

}

func Homehandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, world!")
}