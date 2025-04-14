package services

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/saysb/url-shortener/internal/dto"
	"github.com/saysb/url-shortener/internal/repositories"
)

type ShortenerService struct {
	urlRepo repositories.URLRepository
}

func NewShortenerService(urlRepo repositories.URLRepository) *ShortenerService {
	return &ShortenerService{urlRepo: urlRepo}
}

func (s *ShortenerService) generateShortCode() (string, error) {
	counter, err := s.urlRepo.GetNextCounter()
	if err != nil {
		return "", fmt.Errorf("erreur lors de la récupération de la séquence: %w", err)
	}
	
	seed := counter
	b := make([]byte, 6)
	for i := range b {
		seed = (seed*1103515245 + 12345) & 0x7fffffff
		b[i] = byte(seed % 256)
	}
	
	encoded := base64.URLEncoding.EncodeToString(b)[:6]
	return encoded, nil
}

func (s *ShortenerService) CreateShortURL(originalURL string) (*dto.URLResponse, error) {
	existingURL, err := s.urlRepo.FindByOriginalURL(originalURL)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la recherche de l'URL existante: %w", err)
	}

	if existingURL != nil {
		return &dto.URLResponse{
			OriginalURL: existingURL.OriginalURL,
			ShortCode:   existingURL.ShortCode,
			CreatedAt:   existingURL.CreatedAt,
			AccessCount: existingURL.AccessCount,
		}, nil
	}

	shortCode, err := s.generateShortCode()
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la génération du code court: %w", err)
	}

	newURL, err := s.urlRepo.Create(originalURL, shortCode)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'insertion de l'URL: %w", err)
	}

	return &dto.URLResponse{
		OriginalURL: newURL.OriginalURL,
		ShortCode:   newURL.ShortCode,
		CreatedAt:   newURL.CreatedAt,
		AccessCount: newURL.AccessCount,
	}, nil
}

func (s *ShortenerService) GetOriginalURL(shortCode string) (string, error) {
	url, err := s.urlRepo.FindByShortCode(shortCode)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la recherche de l'URL: %w", err)
	}
	if url == nil {
		return "", fmt.Errorf("URL non trouvée")
	}

	err = s.urlRepo.UpdateAccess(shortCode, time.Now())
	if err != nil {
		return "", fmt.Errorf("erreur lors de la mise à jour des statistiques: %w", err)
	}

	return url.OriginalURL, nil
}

func (s *ShortenerService) DeleteShortURL(shortCode string) error {
	err := s.urlRepo.Delete(shortCode)
	if err != nil {
		return fmt.Errorf("erreur lors de la suppression de l'URL: %w", err)
	}
	return nil
}
