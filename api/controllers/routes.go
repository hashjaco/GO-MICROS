package controllers

import "github.com/hashjaco/GO-MICROS/api/middlewares"

func (s *Server) initRoutes() {

	// Home
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareAuthentication(s.Login)).Methods("POST")

	// User Routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetAllUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
}
