package handler

import (
	"net/url"
	"strings"
	"url-shortener/internal/dto"
	"url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	Service *service.URLService
}

func (h URLHandler) ShortenURL(c *gin.Context) {
	var urls dto.URLRequest

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

	// validate input 
	ori := urls.OriginalUrl

	if ori == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"message": "URL cannot be empty",
			"data": gin.H{
				"original_url": ori,
			},
		})

		return
	}

	// add default scheme if the url didn't have 
	isContainScheme1 := strings.HasPrefix(ori, "https://")
	isContainScheme2 := strings.HasPrefix(ori, "http://")

	if !isContainScheme1 && !isContainScheme2 {
		ori = "https://" + ori
	}

	isValidURL, err := url.ParseRequestURI(ori)

	if err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"message": "Invalid URL Format",
			"data": gin.H{
				"original_url": ori,
			},
		})

		return
	}
	
	if isValidURL.Scheme != "http" && isValidURL.Scheme != "https" {
		c.JSON(400, gin.H{
			"code": 400,
			"message": "Invalid URL Scheme",
			"data": gin.H{
				"scheme": isValidURL.Scheme,
			},
		})

		return
	}

	if isValidURL.Host == "" {
		c.JSON(400, gin.H{
			"code": 400,
			"message": "Invalid URL Host",
			"data": gin.H{
				"host": isValidURL.Host,
			},
		})

		return
	}

	// call services 
	baseURL := "http://localhost:8000"
	shorten, err := h.Service.CreateShortURL(urls.OriginalUrl)

	if err != nil {
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