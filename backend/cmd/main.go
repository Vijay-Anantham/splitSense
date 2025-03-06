package main

import (
	"fmt"
	"splisense/cmd/api"
	"splisense/common/config"
)

func main() {
	// initi basic server config
	configs := config.InitConfig()

	// Initialize DB connection

	server := api.NewApiServer(fmt.Sprintf(":%s", configs.Port), db)
}

func CheckDBConn(db) {

}
