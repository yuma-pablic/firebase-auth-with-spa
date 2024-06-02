package rdb

import (
	"api/ctxx"
	"context"
	"fmt"
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

func (ur *userRepository) Login(ctx context.Context, app *firebase.App) *auth.UserRecord {
	sessionID := ctxx.GetSessions(ctx).ID
	// no set session
	if sessionID == "" {
		return nil
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	const expiresIn = 5 * 60 * 1000
	token, err := client.VerifySessionCookie(ctx, sessionID)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}
	a, err := client.SessionCookie(ctx, token.UID, expiresIn)
	if err != nil {
		log.Fatalf("error creating session cookie: %v\n", err)
	}
	fmt.Println(a)
	return nil
}

func (ur *userRepository) Logout(ctx context.Context, app *firebase.App) error {
	sessionID := ctxx.GetSessions(ctx).ID
	// no set session
	if sessionID == "" {
		return nil
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	decodedClaims, err := client.VerifyIDTokenAndCheckRevoked(ctx, sessionID)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}
	err = client.RevokeRefreshTokens(ctx, decodedClaims.UID)
	if err != nil {
		log.Fatalf("error revoking tokens for user: %v\n", err)
		return err
	}
	return nil
}
