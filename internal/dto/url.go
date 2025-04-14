package dto

import "time"

type URLResponse struct {
	OriginalURL  string    `json:"original_url"`
	ShortCode    string    `json:"short_code"`
	CreatedAt    time.Time `json:"created_at"`
	AccessCount  int       `json:"access_count"`
} 