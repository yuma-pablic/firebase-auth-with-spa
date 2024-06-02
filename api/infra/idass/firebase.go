package idass

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// [END admin_import_golang]

// ==================================================================
// https://firebase.google.com/docs/admin/setup
// ==================================================================

func InitializeAppWithRefreshToken() *firebase.App {
	// [START initialize_app_refresh_token_golang]
	opt := option.WithCredentialsFile("path/to/refreshToken.json")
	config := &firebase.Config{ProjectID: "my-project-id"}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// [END initialize_app_refresh_token_golang]

	return app
}
