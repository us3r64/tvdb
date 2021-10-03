package tvdb

import (
	"context"
	"log"
	"os"
	"testing"
)

var (
	testClient *Client
	testAPIKey = os.Getenv("TVDB_APIKEY")
	testPin    = os.Getenv("TVDB_PIN")
)

func TestMain(m *testing.M) {
	testClient = New(testAPIKey, testPin)

	// login
	err := testClient.Login(context.Background())
	if err != nil {
		log.Fatalln(err, "Login error")
	}

	// run tests
	code := m.Run()

	os.Exit(code)
}
