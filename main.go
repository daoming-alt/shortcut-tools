package main

import (
	"fmt"
	"github.com/kdbrian/shortcut-tool/config"
	"log"
	"net/http"
)

func main() {

	err := database.Connect()
	if err != nil {
		log.Fatal("failed to connect db : ", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Shortcuts at /short-x")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
