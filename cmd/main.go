package main

import (
	"fmt"
	"url-shortener/internal/config"
	"url-shortener/internal/repository"
)

func main() {
	// connect with db
	db := config.ConnectDB()

	// init repo
	repo := repository.URLRepository{DB: db}
	
	// init service 
	urlService := Service.CreateShortURL(repo)

	fmt.Println("Running the server at port 8000")
}