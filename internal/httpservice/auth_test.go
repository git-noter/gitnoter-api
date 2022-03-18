package httpservice

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vivekweb2013/batnoter/internal/config"
)

func TestGithubLogin(t *testing.T) {
	t.Run("should redirect to provider when the github login request is valid", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		router := gin.Default()
		handler := NewAuthHandler(nil, nil, config.OAuth2{})

		router.GET("/api/v1/oauth2/login/github", handler.GithubLogin)
		response := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/oauth2/login/github", nil)

		router.ServeHTTP(response, req)
		assert.Equal(t, http.StatusTemporaryRedirect, response.Code)
	})
}
