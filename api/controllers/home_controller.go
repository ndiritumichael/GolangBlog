package controllers

import (
	"net/http"

	"github.com/ndiritumichael/goblog/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {

	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
	responses.JSON(w, http.StatusOK, "to view users  use '/users'")

}
