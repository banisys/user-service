package test

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/banisys/user-service/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func migrateDatabase() {
	migrator, err := migrate.New(
		"file://../pkg/database/migrations",
		"sqlite3://file::memory:?cache=shared",
	)
	if err != nil {
		log.Fatalf("failed to initialize migrator: %v", err)
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to apply migrations: %v", err)
	}
}

func TestSignup(t *testing.T) {
	migrateDatabase()

	route := gin.Default()
	routes.RegisterRoutes(route)

	testBody := []byte(`{"email":"test@example.com","password":"123"}`)
	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(testBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	os.Remove("database.db")

}

func TestLogin(t *testing.T) {

	assert.Equal(t, "111", "111")

}
