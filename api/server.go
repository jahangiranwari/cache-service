package api

import (
	"log"
	"net/http"
	"os"
)

func StartServer() {
	port := os.Getenv("PORT")
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/github/", gitHubHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
