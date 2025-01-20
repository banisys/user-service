package test

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)
	setupTestEnvironment()

	exitCode := m.Run()
	teardown()

	os.Exit(exitCode)
}

func setupTestEnvironment() {
	os.Setenv("ENV", "test")
}

func teardown() {
	os.Unsetenv("ENV")
}
