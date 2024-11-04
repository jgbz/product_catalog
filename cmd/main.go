package main

import (
	"github.com/jgbz/product_catalog/database"
	"github.com/jgbz/product_catalog/server"
	"github.com/jgbz/product_catalog/service"
)

func main() {
	repo, err := database.NewRepository()
	if err != nil {
		panic(err)
	}
	if err := repo.RunMigration(); err != nil {
		panic(err)
	}

	service := service.NewFeedService(repo)

	server := server.New(service)
	server.Listen()
}
