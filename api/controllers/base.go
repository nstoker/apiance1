package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// Server server
type Server struct {
	DB     *sqlx.DB
	Router *mux.Router
}

// InitializeDatabase initializes the database
func (server *Server) InitializeDatabase(databaseURI string) error {

	var err error

	db, err := sqlx.Open("postgres", databaseURI)
	if err != nil {
		return fmt.Errorf("Server:Initialize error opening: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("Server:Initialize error pinging %w", err)
	}

	server.DB = db

	return nil
}

// InitializeRouter initialises the router
func (server *Server) InitializeRouter() error {

	server.Router = mux.NewRouter()

	server.initializeRoutes()

	return nil
}

// Run server, run. See server run
func (server *Server) Run(addr string) error {
	fmt.Println("Listening to port 8080")
	return fmt.Errorf("base.Run(): %w", http.ListenAndServe(addr, server.Router))
}
