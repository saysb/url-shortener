package repositories

import (
	"database/sql"
	"time"

	"github.com/saysb/url-shortener/internal/models"
)

type PostgresURLRepository struct {
	db *sql.DB
}

func NewPostgresURLRepository(db *sql.DB) *PostgresURLRepository {
	return &PostgresURLRepository{db: db}
}

func (r *PostgresURLRepository) FindByOriginalURL(originalURL string) (*models.URL, error) {
	var url models.URL
	err := r.db.QueryRow(
		"SELECT id, original_url, short_code, created_at, last_accessed, access_count FROM urls WHERE original_url = $1",
		originalURL,
	).Scan(&url.ID, &url.OriginalURL, &url.ShortCode, &url.CreatedAt, &url.LastAccessed, &url.AccessCount)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *PostgresURLRepository) FindByShortCode(shortCode string) (*models.URL, error) {
	var url models.URL
	err := r.db.QueryRow(
		"SELECT id, original_url, short_code, created_at, last_accessed, access_count FROM urls WHERE short_code = $1",
		shortCode,
	).Scan(&url.ID, &url.OriginalURL, &url.ShortCode, &url.CreatedAt, &url.LastAccessed, &url.AccessCount)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *PostgresURLRepository) Create(originalURL, shortCode string) (*models.URL, error) {
	var url models.URL
	err := r.db.QueryRow(
		"INSERT INTO urls (original_url, short_code) VALUES ($1, $2) RETURNING id, original_url, short_code, created_at, access_count",
		originalURL, shortCode,
	).Scan(&url.ID, &url.OriginalURL, &url.ShortCode, &url.CreatedAt, &url.AccessCount)

	if err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *PostgresURLRepository) UpdateAccess(shortCode string, accessTime time.Time) error {
	_, err := r.db.Exec(
		"UPDATE urls SET access_count = access_count + 1, last_accessed = $1 WHERE short_code = $2",
		accessTime, shortCode,
	)
	return err
}

func (r *PostgresURLRepository) Delete(shortCode string) error {
	_, err := r.db.Exec("DELETE FROM urls WHERE short_code = $1", shortCode)
	return err
}

func (r *PostgresURLRepository) GetNextCounter() (int64, error) {
	var counter int64
	err := r.db.QueryRow("SELECT nextval('urls_counter_seq')").Scan(&counter)
	return counter, err
}