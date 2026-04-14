package main

import (
	"fmt"
	"url-shortener/internal/config"
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
)

func main() {
	db := config.ConnectDB()
	url := model.URL{
		OriginalUrl: "https://google.com",
	}

	repo := repository.URLRepository{DB: db}

	createNewData, err := repo.Insert(url)

	fmt.Println(createNewData, err)

	fmt.Println("Running the server at port 8000")
}