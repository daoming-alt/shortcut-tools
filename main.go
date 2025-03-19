package main

import (
	"fmt"
	"github.com/kdbrian/shortcut-tool/config"
	"log"
	"net/http"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env.local")

	err := database.Connect()
	if err != nil {
		log.Fatal("failed to connect db : ", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Shortcuts at /short-x")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
