package repository

import (
	"url-shortener/internal/model"

	"github.com/jmoiron/sqlx"
)

type URLRepository struct {
	DB *sqlx.DB
}

// func Insert (db URLRepository, url *model.URL) int {
// 	var id int 

// 	insert := db.DB.QueryRow("INSERT INTO urls (original_url) VALUES ($1) RETURNING id", url.OriginalUrl).Scan(&id)
	
// 	return data
// }

func (repo URLRepository) Insert(url model.URL) (int, error) {
	var id int
	
	err := repo.DB.QueryRow("INSERT INTO urls (original_url) VALUES ($1) RETURNING id", url.OriginalUrl).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil 
}