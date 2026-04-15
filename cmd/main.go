package main

import (
	"fmt"
	"url-shortener/internal/config"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"
)

func main() {
	db := config.ConnectDB()
	
	// url := model.URL{
	// 	OriginalUrl: "https://www.linkedin.com/in/fajryalvinhidayat/",
	// }
	// createNewData, _ := repo.Insert(url)

	repo := repository.URLRepository{DB: db}

	shortCode :=  service.EncodeToBase62(2)
	updShortCode := repo.UpdateShortCode(2, shortCode)
	fmt.Println(updShortCode)

	fmt.Println("Running the server at port 8000")
}