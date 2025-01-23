package test

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {

	gin.SetMode(gin.TestMode)

	exitCode := m.Run()

	os.Exit(exitCode)
}
