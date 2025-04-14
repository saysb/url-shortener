package models

import (
	"database/sql"
	"time"
)

type URL struct {
    ID           int       `json:"id" db:"id"`
    OriginalURL  string    `json:"original_url" db:"original_url"`
    ShortCode    string    `json:"short_code" db:"short_code"`
    CreatedAt    time.Time `json:"created_at" db:"created_at"`
    LastAccessed sql.NullTime `json:"last_accessed,omitempty" db:"last_accessed"`
    AccessCount  int       `json:"access_count" db:"access_count"`
}

