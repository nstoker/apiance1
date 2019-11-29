package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/nstoker/apiance1/api/migrate"
)

// Server server
type Server struct {
	DB     *sqlx.DB
	Router *mux.Router
}

// Initialize initialize
func (server *Server) Initialize(databaseURI string) error {

	var err error

	db, err := sqlx.Open("postgres", databaseURI)
	if err != nil {
		return fmt.Errorf("Server:Initialize error opening: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("Server:Initialize error pinging %w", err)
	}

	if err := migrate.Perform(); err != nil {
		return fmt.Errorf("Server:Initialize error migrating: %w", err)
	}

	server.Router = mux.NewRouter()

	server.initializeRoutes()

	return nil
}

// Run server, run. See server run
func (server *Server) Run(addr string) error {
	fmt.Println("Listening to port 8080")
	return fmt.Errorf("base.Run(): %w", http.ListenAndServe(addr, server.Router))
}
