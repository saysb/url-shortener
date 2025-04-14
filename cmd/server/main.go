package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/saysb/url-shortener/internal/api"
	"github.com/saysb/url-shortener/internal/repositories"
	"github.com/saysb/url-shortener/internal/services"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName))
	if err != nil {
		log.Fatalf("Erreur lors de la connexion à la base de données: %v", err)
	}
	defer db.Close()

	urlRepo := repositories.NewPostgresURLRepository(db)

	shortenerService := services.NewShortenerService(urlRepo)

	r := api.SetupRouter(shortenerService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Serveur démarré sur le port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

