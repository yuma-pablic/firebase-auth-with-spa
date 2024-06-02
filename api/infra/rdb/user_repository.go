package rdb

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

func getUser(ctx context.Context, app *firebase.App) *auth.UserRecord {
	uid := "some_string_uid"

	// [START get_user_golang]
	// Get an auth client from the firebase.App
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	u, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	log.Printf("Successfully fetched user data: %v\n", u)
	// [END get_user_golang]
	return u
}

func deleteUser(ctx context.Context, client *auth.Client) {
	uid := "d"
	// [START delete_user_golang]
	err := client.DeleteUser(ctx, uid)
	if err != nil {
		log.Fatalf("error deleting user: %v\n", err)
	}
	log.Printf("Successfully deleted user: %s\n", uid)
	// [END delete_user_golang]
}

func logoutUser(ctx context.Context, client *auth.Client, uid string) {
	// [START revoke_tokens_golang]
	if err := client.RevokeRefreshTokens(ctx, uid); err != nil {
		log.Fatalf("error revoking tokens for user: %v, %v\n", uid, err)
	}
	// [END revoke_tokens_golang]
}
