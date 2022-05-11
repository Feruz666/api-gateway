package pkg

import (
	"github.com/Feruz666/api-gateway/pkg/auth/handlers"
	"github.com/Feruz666/api-gateway/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests and responses
type Server struct {
	config util.Config
	router *gin.Engine
}

func ServiceRoutes(config util.Config) (*Server, error) {
	server := &Server{
		config: config,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/auth", handlers.CreateUser)
	router.POST("/auth/login", handlers.LoginUser)

	server.router = router
}

// Start runs server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
