package rdb

import (
	"api/ctxx"
	"context"
	"fmt"
	"log"

	userDomain "api/domain/user"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type userRepository struct {
	firebase *firebase.App
}

func NewUserRepository(firebase *firebase.App) *userRepository {
	return &userRepository{firebase}
}

func (ur *userRepository) Find(ctx context.Context, app *firebase.App) (*userDomain.User, error) {
	sessionID := ctxx.GetSessions(ctx).ID
	// no set session
	if sessionID == "" {
		return nil, fmt.Errorf("no session")
	}
	csrfToken := ctxx.GetCSRFToken(ctx)
	if csrfToken == "" {
		return nil, fmt.Errorf("no csrf token")
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}
	token, err := client.VerifySessionCookie(ctx, sessionID)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v", err)
	}
	res, err := client.GetUser(ctx, token.UID)
	if err != nil {
		log.Fatalf("error getting user: %v\n", err)
		return nil, err
	}
	user := &userDomain.User{
		ID:    res.UID,
		Email: res.Email,
	}
	return user, nil
}

func (ur *userRepository) Login(ctx context.Context, app *firebase.App) (*auth.UserRecord, error) {
	sessionID := ctxx.GetSessions(ctx).ID
	// no set session
	if sessionID == "" {
		return nil, nil
	}
	csrfToken := ctxx.GetCSRFToken(ctx)
	if csrfToken == "" {
		return nil, nil
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		return nil, err
	}
	const expiresIn = 5 * 60 * 1000
	token, err := client.VerifySessionCookie(ctx, sessionID)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}
	a, err := client.SessionCookie(ctx, token.UID, expiresIn)
	if err != nil {
		log.Fatalf("error creating session cookie: %v\n", err)
		return nil, err
	}
	fmt.Println(a)
	return nil, err
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
		return err
	}
	decodedClaims, err := client.VerifyIDTokenAndCheckRevoked(ctx, sessionID)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
		return err
	}
	err = client.RevokeRefreshTokens(ctx, decodedClaims.UID)
	if err != nil {
		log.Fatalf("error revoking tokens for user: %v\n", err)
		return err
	}
	return nil
}

func (ur *userRepository) Delete(ctx context.Context, app *firebase.App) error {
	sessionID := ctxx.GetSessions(ctx).ID
	// no set session
	if sessionID == "" {
		return nil
	}
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
		return err
	}
	decodedClaims, err := client.VerifySessionCookie(ctx, sessionID)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
		return err
	}
	err = client.DeleteUser(ctx, decodedClaims.UID)
	if err != nil {
		log.Fatalf("error deleting user: %)v\n", err)
		return err
	}
	return nil
}
