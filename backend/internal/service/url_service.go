package service

import (
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
)

type URLService struct {
	Repo *repository.URLRepository
}

func EncodeToBase62(id int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// convert id by encoded string
	if id == 0 {
		return string(charset[id])
	}

	var encoded string

	for id > 0 {
		remainder := id % 62
		encoded += string(charset[remainder])
		id = id / 62
	}

	var reversed string

	// reverse the encoded string
	for j := len(encoded) - 1; j >= 0; j-- {
		reversed += string(encoded[j])
	}

	return reversed
}

func (s URLService) CreateShortURL(original_url string) (string, error) {

	// panggil repo.insert => dapet id
	createdId, err := s.Repo.Insert(model.URL{OriginalUrl: original_url})

	if err != nil {
		return "", err
	}

	// generate short_code dengan cara encode retrieved id 
	shortCode := EncodeToBase62(createdId)
	 
	// panggil repo.updateShortCode 
	err = s.Repo.UpdateShortCode(createdId, shortCode)

	if err != nil {
		return "", err
	}

	return shortCode, nil
}

func (s URLService) GetOriginalURL(short_code string) (string, error) {
	origin, err := s.Repo.FindByShortCode(short_code)

	if err != nil {
		return "", err
	}

	return origin, nil
}