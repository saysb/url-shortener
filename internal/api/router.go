package api

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/saysb/url-shortener/internal/services"
)

type shortenRequest struct {
	URL string `json:"url"`
}

func apiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		validApiKey := os.Getenv("API_KEY")

		if apiKey == "" || apiKey != validApiKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SetupRouter(service *services.ShortenerService) *chi.Mux {
    r := chi.NewRouter()
    
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)
    
    r.Route("/api/v1", func(r chi.Router) {
        r.Use(apiKeyMiddleware)
        
        r.Post("/shorten", func(w http.ResponseWriter, r *http.Request) {
            var req shortenRequest
            if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
                http.Error(w, "Invalid request body", http.StatusBadRequest)
                return
            }

            if req.URL == "" {
                http.Error(w, "URL is required", http.StatusBadRequest)
                return
            }

            url, err := service.CreateShortURL(req.URL)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(url)
        })

        r.Delete("/{code}", func(w http.ResponseWriter, r *http.Request) {
            code := chi.URLParam(r, "code")
            if code == "" {
                http.Error(w, "Code is required", http.StatusBadRequest)
                return
            }
            err := service.DeleteShortURL(code)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
            }
        })
    })
    
    r.Get("/{code}", func(w http.ResponseWriter, r *http.Request) {
        code := chi.URLParam(r, "code")
        originalURL, err := service.GetOriginalURL(code)
        if err != nil {
            tmplPath := "/app/internal/templates/404.html"
        
            tmpl, err := template.ParseFiles(tmplPath)
            if err != nil {
                http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
                return
            }
            
            w.WriteHeader(http.StatusNotFound)
            if err := tmpl.Execute(w, nil); err != nil {
                http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
                return
            }
            return
        }
        http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
    })
    
    return r
}
