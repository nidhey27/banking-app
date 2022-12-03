package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nidhey27/banking-app/db/sql"
)

// Server serves to HTTP Request for our banking service
type Server struct {
	store  sql.Store
	router *gin.Engine
}

// NewServer creates a new Http server and setup routing
func NewServer(store sql.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/", server.listAccounts)

	router.PUT("/accounts/:id", server.updateAccount)

	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router

	return server
}

// Start runs the HTTP server on Specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
