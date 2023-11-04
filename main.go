package main

import (
	"app/config"
	"app/container"
	"app/pkg/database"
	"app/server"
	"fmt"
)

func init() {
	config.NewConfig()
}

func main() {
	database.NewMysqlDB()
	r := server.NewRestapi()
	container.NewContainer(r)

	err := r.Run(fmt.Sprintf(":%s", config.App.Port))
	if err != nil {
		panic("routing error")
	}
}
