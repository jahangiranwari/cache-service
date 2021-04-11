package api

import (
	"log"
	"net/http"
	"os"
)

func StartServer() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
