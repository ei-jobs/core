package main

import (
	"fmt"

	server "github.com/aidosgal/gust/internal/app"
	"github.com/aidosgal/gust/internal/config"
	"github.com/aidosgal/gust/internal/database"
)

func main() {
	cfg := config.MustLoad()

	db, err := database.NewDatabase(cfg.Database)
	if err != nil {
		panic(err)
	}

	server := server.NewAPIServer(fmt.Sprintf("localhost:%d", cfg.Server.Port), db)

	err = server.Run()
	if err != nil {
		panic(err)
	}
}
