package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {

	fmt.Println("11111111")

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
