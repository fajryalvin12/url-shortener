package handler

import (
	"url-shortener/internal/model"
	"url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	Service *service.URLService
}

func (h URLHandler) ShortenURL(c *gin.Context) {
	var urls model.URL

	// retrieve data from json request body, then bind it 
	err := c.ShouldBindJSON(&urls) 
	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"message": "Failed to bind request",
			"error": err.Error(),
		})

		return
	} 

	// call services 
	baseURL := "http://localhost:8000"
	shorten, err := h.Service.CreateShortURL(urls.OriginalUrl)

	if err !=  nil {
		c.JSON(500, gin.H{
			"code": 500,
			"message": "Failed to create shorten url",
			"error": err.Error(),
		})

		return
	}

	shortURL := baseURL + "/" + shorten


	c.JSON(201, gin.H{
		"code": 201,
		"message": "URL shortened successfully",
		"data": gin.H{
			"short_url": shortURL,
			"original_url": urls.OriginalUrl,
		},
	})
}

func (h URLHandler) RedirectURL(c *gin.Context) {
	shortCode := c.Param("short_code")

	originalURL, err := h.Service.GetOriginalURL(shortCode) 
	if err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"message": "URL Not found",
			"error": err.Error(),
		})

		return
	}

	c.Redirect(302, originalURL)
}