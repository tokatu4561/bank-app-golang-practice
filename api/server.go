package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/tokatu4561/simple-bank/db/sqlc"
)

type Server struct {
	store db.Store
	router  *gin.Engine
}

func NewSever(store db.Store) (*Server, error) {
	server := &Server{
		store: store,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}