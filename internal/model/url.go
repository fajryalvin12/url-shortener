package model

import "time"

type URL struct {
	Id          int64   	`db:"id"`
	OriginalUrl string 		`json:"original_url" db:"original_url"`
	ShortCode   string 		`db:"short_code"`
	CreatedAt   time.Time 	`db:"created_at"`
}