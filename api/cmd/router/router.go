package router

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
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
	opt := option.WithCredentialsFile("path/to/service-account-file.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	ctx := r.Context()
	client, err := app.Auth(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	token, err := client.VerifyIDToken(ctx, "token")
	users := []string{"Alice", "Bob", "Charlie"}
	json.NewEncoder(w).Encode(users)
}

func UsersLoginHandler(w http.ResponseWriter, r *http.Request) {
	opt := option.WithCredentialsFile("path/to/service-account-file.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	ctx := r.Context()
	// Get the ID token sent by the client
	defer r.Body.Close()
	idToken, err := getIDTokenFromBody(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set session expiration to 5 days.
	expiresIn := time.Hour * 24 * 5

	// Create the session cookie. This will also verify the ID token in the process.
	// The session cookie will have the same claims as the ID token.
	// To only allow session cookie setting on recent sign-in, auth_time in ID token
	// can be checked to ensure user was recently signed in before creating a session cookie.
	cookie, err := client.SessionCookie(r.Context(), idToken, expiresIn)
	if err != nil {
		http.Error(w, "Failed to create a session cookie", http.StatusInternalServerError)
		return
	}

	// Set cookie policy for session cookie.
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    cookie,
		MaxAge:   int(expiresIn.Seconds()),
		HttpOnly: true,
		Secure:   true,
	})
	w.Write([]byte(`{"status": "success"}`))
}

func UsersLogoutHandler(w http.ResponseWriter, r *http.Request) {

	// Set cookie policy for session cookie.
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    cookie,
		MaxAge:   int(expiresIn.Seconds()),
		HttpOnly: true,
		Secure:   true,
	})
	w.Write([]byte(`{"status": "success"}`))
}
