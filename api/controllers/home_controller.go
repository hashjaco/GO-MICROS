package controllers

import(
	"github.com/hashjaco/GO-MICROS/api/responses"
	"net/http"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "WELCOME TO THE GO-LANG API MICROSERVICE")
}
