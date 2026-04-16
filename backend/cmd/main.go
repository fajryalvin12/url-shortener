package main

import (
	"url-shortener/internal/config"
	"url-shortener/internal/handler"
	"url-shortener/internal/repository"
	"url-shortener/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// DB
	db := config.ConnectDB()

	// repo
	repo := repository.URLRepository{DB: db}

	// service
	urlService := service.URLService{&repo}

	// handler
	urlHandler := handler.URLHandler{Service: &urlService}

	// router + cors middleware assignment
	r := gin.Default()

	r.Use(cors.Default())

	// routes
	r.POST("/v1/shorten", urlHandler.ShortenURL)
	r.GET("/:short_code", urlHandler.RedirectURL)

	// run
	r.Run(":8000")
}