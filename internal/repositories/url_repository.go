package repositories

import (
	"time"

	"github.com/saysb/url-shortener/internal/models"
)

type URLRepository interface {
	FindByOriginalURL(originalURL string) (*models.URL, error)
	FindByShortCode(shortCode string) (*models.URL, error)
	Create(originalURL, shortCode string) (*models.URL, error)
	UpdateAccess(shortCode string, accessTime time.Time) error
	Delete(shortCode string) error
	GetNextCounter() (int64, error)
}