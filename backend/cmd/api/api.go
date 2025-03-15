package api

import (
	"database/sql"
	"log"
	userhandler "splisense/services/userHandler"

	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	router := gin.Default()

	v1 := router.Group("/v1")

	userService := userhandler.NewHandler()
	userService.RegisterRouts(v1)

	log.Println("Listening on", s.addr)

	return router.Run(s.addr)
}
