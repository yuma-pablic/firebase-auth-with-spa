package rdb

import (
	"api/ctxx"
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type userRepository struct {
	firebase *firebase.App
}

func NewUserRepository(firebase *firebase.App) *userRepository {
	return &userRepository{firebase}
}

func (ur *userRepository) Find(ctx context.Context, app *firebase.App) *auth.UserRecord {
	sessionID := ctxx.GetSessions(ctx).ID
	// no set session
	if sessionID == "" {
		return nil
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	token, err := client.VerifySessionCookie(ctx, sessionID)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}
	u, err := client.GetUser(ctx, token.UID)
	if err != nil {
		log.Fatalf("error getting user: %v\n", err)
	}
	return u
}

func (ur *userRepository) Create(ctx context.Context, app *firebase.App, email string, password string) *auth.UserRecord {
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password)
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	return u
}
