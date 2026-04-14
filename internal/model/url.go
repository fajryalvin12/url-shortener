package model

import "time"

type URL struct {
	Id          int64   	`db:"id"`
	OriginalUrl string 		`db:"original_url"`
	ShortCode   string 		`db:"short_code"`
	CreatedAt   time.Time 	`db:"created_at"`
}