package middlewares

import (
	"dog/api/server"
	"dog/static"
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
	"path"
)

func AddDocumentationMiddleware(server *server.Server) {
	opts := middleware.RedocOpts{SpecURL: "/static/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	server.Router.Handle("GET", "/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Dog Version 0.0.1",
		})
	})

	server.Router.GET("/static/*filepath", func(c *gin.Context) {
		_, fileName := path.Split(c.Request.URL.Path)
		c.FileFromFS(fileName, http.FS(static.Files))
	})
	server.Router.Handle("GET", "/docs", gin.WrapH(sh))
}
