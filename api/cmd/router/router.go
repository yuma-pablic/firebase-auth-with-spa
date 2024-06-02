package router

import (
	"encoding/json"
	"net/http"
)

func Run() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	http.HandleFunc("/api/users/me", UsersHandler)
	http.HandleFunc("/fuga", UsersHandler)
	http.HandleFunc("/api/login", UsersLoginHandler)
	server.ListenAndServe()
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []string{"Alice", "Bob", "Charlie"}
	json.NewEncoder(w).Encode(users)
}

func UsersLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status": "success"}`))
}

func UsersLogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status": "success"}`))
}
