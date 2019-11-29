package controllers

import (
	"net/http"

	"github.com/nstoker/apiance1/api/responses"
)

// Home the home server
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}
