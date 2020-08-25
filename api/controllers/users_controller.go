package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hashjaco/GO-MICROS/api/models"
	"github.com/hashjaco/GO-MICROS/api/responses"
	"github.com/hashjaco/GO-MICROS/api/utils/format_error"
	"io/ioutil"
	"net/http"
	"strconv"
)

/**
* Handle GET requests to return all users
**/
func (s *Server) GetAllUsers(w http.ResponseWriter, r *http.Request){

	user := models.User{}

	users, err := user.FindAllUsers(s.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}


/**
* Handle GET requests with ID parameter
**/
func (s *Server) GetUser(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userRetrieved, err := user.FindUserByID(s.DB, uid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userRetrieved)
}


/**
* Handle POST requests
**/
func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(s.DB)

	if err != nil {

		formattedError := format_error.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
}




