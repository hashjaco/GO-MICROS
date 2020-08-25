package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hashjaco/GO-MICROS/api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (s *Server) Init(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	if DbDriver == "mysql" {
		DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		s.DB, err = gorm.Open(DbDriver, DbUrl)
		if err != nil {
			fmt.Printf("cannot connect to %s database", DbDriver)
			log.Fatal("This is the error: ", err)
		} else {
			fmt.Printf("successfully connected to %s database", DbDriver)
		}
	}

	s.DB.Debug().AutoMigrate(&models.User{})

	s.Router = mux.NewRouter()

	s.initRoutes()
}

func (s *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, s.Router))
}