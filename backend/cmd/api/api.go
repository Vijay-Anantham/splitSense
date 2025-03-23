package api

import (
	"database/sql"
	"log"
	expensehandler "splisense/services/expenseHandler"
	grouphandler "splisense/services/groupHandler"
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

	userStore := userhandler.NewStore(s.db)
	userService := userhandler.NewHandler(userStore)
	userService.RegisterRouts(v1)

	groupStore := grouphandler.NewStore(s.db)
	groupService := grouphandler.NewHandler(groupStore)
	groupService.RegisterRoutes(v1)

	expenseStore := expensehandler.NewStore(s.db)
	expenseService := expensehandler.NewHandler(expenseStore)
	expenseService.RegisterRoutes(v1)

	log.Println("Listening on", s.addr)

	return router.Run(s.addr)
}
