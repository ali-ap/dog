package server

import (
	"dog/api/v1/controllers"
	"dog/configs"
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	Router *gin.Engine
}

func NewServer(path string) *Server {
	ReadConfig(path)
	router := gin.Default()
	router.Use(cors.Default())
	RegisterRoutes(router)
	initializeServer()
	return &Server{
		Router: router,
	}
}

var RegisterRoutes = func(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/dig/", controllers.Dig)
	}
}

var initializeServer = func() {
	myFigure := figure.NewColorFigure("DOG", "puffy", "purple", true)
	myFigure.Print()
}

func (server *Server) Run(port string) {
	err := server.Router.Run(fmt.Sprintf(":%v", port))
	if err != nil {
		panic(fmt.Sprintf("can not run the router error : %v", err))
	}
}

func ReadConfig(path string) {
	config, err := configs.NewConfig(path)
	if err != nil {
		log.Fatal(err)
	}
	configs.AppConfig = config
}
