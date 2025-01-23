package test

import (
	"os"
	"testing"

	"github.com/banisys/user-service/pkg/utils"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	utils.LoadConfig(".")

	gin.SetMode(gin.TestMode)

	exitCode := m.Run()

	os.Exit(exitCode)
}
