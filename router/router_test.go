package router

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestSetupRouterReleaseModeWhenDebugOff(t *testing.T) {
	os.Unsetenv("DEBUG")

	SetupRouter()

	assert.Equal(t, gin.Mode(), gin.ReleaseMode)
}

func TestSetupRouterDebugModeWhenDebugOn(t *testing.T) {
	os.Setenv("DEBUG", "true")
	defer os.Unsetenv("DEBUG")

	SetupRouter()

	assert.Equal(t, gin.Mode(), gin.DebugMode)
}

func TestSetupRouterHasExpectedRoutes(t *testing.T) {
	os.Unsetenv("DEBUG")

	r := SetupRouter()

	routes := r.Routes()
	routeMap := make(map[string]bool)
	for _, route := range routes {
		routeMap[route.Method+":"+route.Path] = true
	}

	assert.Equal(t, routeMap["POST:/message"], true)
	assert.Equal(t, routeMap["GET:/message"], true)
	assert.Equal(t, routeMap["POST:/directedMessage"], true)
	assert.Equal(t, routeMap["POST:/subscribe/message"], true)
	assert.Equal(t, routeMap["POST:/unsubscribe/message"], true)
}
