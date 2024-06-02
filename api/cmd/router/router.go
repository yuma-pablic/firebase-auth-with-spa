package router

import (
	"api/controller/user"
	"net/http"

	"github.com/gorilla/sessions"
)

func Run() {

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	http.HandleFunc("/api/users/me", user.UserHandler.Get)
	http.HandleFunc("/api/login", UsersLoginHandler)
	http.HandleFunc("/api/delete/logout", UsersLogoutHandler)
	server.ListenAndServe()
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
