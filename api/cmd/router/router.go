package router

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
)

func Run() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	http.HandleFunc("/api/users/me", UsersHandler)
	http.HandleFunc("/fuga", UsersHandler)
	http.HandleFunc("/api/login", UsersLoginHandler)
	http.HandleFunc("/api/delete/logout", UsersLogoutHandler)
	server.ListenAndServe()
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		if errors.As(err, &err) {
			http.Error(w, err.Error(), http.StatusNonAuthoritativeInfo)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	users := []string{"user1", "user2", "user3"}
	json.NewEncoder(w).Encode(users)
}

func UsersLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status": "success"}`))
}

func UsersLogoutHandler(w http.ResponseWriter, r *http.Request) {
	clearCookie(w, r)
	w.Write([]byte(`{"status": "success"}`))
}

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func clearCookie(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	// セッションの値をクリアする
	session.Values = make(map[interface{}]interface{})

	// セッションを保存する
	session.Save(r, w)

	// クライアントにメッセージを送信する
	w.Write([]byte("Session cleared"))
}
