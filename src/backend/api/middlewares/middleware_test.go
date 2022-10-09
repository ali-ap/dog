package middlewares

import (
	"dog/api/server"
	"flag"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type TestDouble struct {
	Server *server.Server
}

var Mock *TestDouble

func TestMain(m *testing.M) {
	Mock = new(TestDouble)
	Mock.Server = server.NewServer("../../configs/app_config.yml")
	AddDocumentationMiddleware(Mock.Server)
	flag.Parse()
	exitCode := m.Run()
	os.Exit(exitCode)
}
func TestAddDocumentationMiddleware(t *testing.T) {
	t.Run("get application version from index", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		Mock.Server.Router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "{\"message\":\"Dog Version 0.0.1\"}", w.Body.String())
	})

	t.Run("get application open api documentation", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/docs", nil)
		w := httptest.NewRecorder()
		Mock.Server.Router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		assert.NotEmpty(t, w.Body.String())
	})
}
