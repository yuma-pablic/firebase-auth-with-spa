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
	// [START get_user_golang]
	// Get an auth client from the firebase.App
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	u, err := client.GetUser(ctx, sessionID)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", sessionID, err)
	}
	log.Printf("Successfully fetched user data: %v\n", u)
	// [END get_user_golang]
	return u
}
