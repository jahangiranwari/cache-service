package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jahangiranwari/cache-service/cache"
	"github.com/jahangiranwari/cache-service/httputil"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome!!"))
}

func gitHubHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/github/"):]
	value := cache.Query(key)
	if len(value) == 0 {
		value = httputil.GetGitHubRepo(key)
		// Simulate slow API response
		time.Sleep(3 * time.Second)
		cache.Save(key, value)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json.RawMessage(value))
}

func clearCacheHandler(w http.ResponseWriter, r *http.Request) {
	cache.FlushDB()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Cache cleared"))
}

func webhooksHandler(w http.ResponseWriter, r *http.Request) {
	signature := r.Header.Get("x-hub-signature")
	body, _ := ioutil.ReadAll(r.Body)

	if !httputil.VerifySignature(signature, body) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid signature!"))
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - Done!"))
	}
}
