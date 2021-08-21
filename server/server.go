package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joelgarciajr84/go-rest-api/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "3000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Print("SERVER IS RUNNING ON PORT: " + s.port)
	log.Fatal(router.Run(":" + s.port))
}
