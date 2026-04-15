package repository

import (
	"url-shortener/internal/model"

	"github.com/jmoiron/sqlx"
)

type URLRepository struct {
	DB *sqlx.DB
}

func (repo URLRepository) Insert(url model.URL) (int, error) {
	var id int
	
	err := repo.DB.QueryRow("INSERT INTO urls (original_url) VALUES ($1) RETURNING id", url.OriginalUrl).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil 
}

func (repo URLRepository) UpdateShortCode (id int, shortCode string) error {
	_, err := repo.DB.Exec("UPDATE urls SET short_code = $1 WHERE id = $2", shortCode, id)

	if err != nil {
		return err
	}

	return nil
}