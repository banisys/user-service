package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/banisys/user-service/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSignupRoute(t *testing.T) {
	// Initialize the router from the main application
	route := gin.Default()
	routes.RegisterRoutes(route)

	// Prepare test data
	testBody := []byte(`{"email":"test@example.com","password":"123"}`)
	req, err := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(testBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	recorder := httptest.NewRecorder()
	route.ServeHTTP(recorder, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, recorder.Code) // Check status code
}
