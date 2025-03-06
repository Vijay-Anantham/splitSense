package api

import (
	"database/sql"
	"log"

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

	log.Println("Listening on", s.addr)

	return router.Run(s.addr)
}
