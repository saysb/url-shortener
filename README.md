# URL Shortener

Un service de raccourcissement d'URL simple et efficace, développé en Go.

## Fonctionnalités

- Raccourcissement d'URL
- Redirection automatique
- API RESTful
- Base de données PostgreSQL
- Support Docker

## Prérequis

- Go 1.23.5 ou supérieur
- PostgreSQL
- Docker et Docker Compose (optionnel)

## Installation

### Configuration locale

1. Clonez le dépôt :
```bash
git clone https://github.com/saysb/url-shortener.git
cd url-shortener
```

2. Créez un fichier `.env` à la racine du projet avec les variables suivantes :
```env
DB_NAME=votre_nom_de_db
DB_USER=votre_utilisateur
DB_PASSWORD=votre_mot_de_passe
DB_HOST=localhost
DB_PORT=5432
API_KEY=votre_clé_api
```

3. Initialisez la base de données :
```bash
make db-reset
make migrate-up
```

4. Lancez l'application :
```bash
make run
```

### Utilisation avec Docker

1. Créez un fichier `.env` à la racine du projet avec les variables suivantes :
```env
DB_NAME=votre_nom_de_db
DB_USER=votre_utilisateur
DB_PASSWORD=votre_mot_de_passe
DB_HOST=db-url-shortener
DB_PORT=5433
API_KEY=votre_clé_api
```

2. Lancez les conteneurs :
```bash
docker-compose up -d
```

L'application sera accessible sur `http://localhost:8080`

## Commandes disponibles

- `make run` : Lance l'application
- `make migrate-up` : Applique les migrations
- `make migrate-down` : Annule les migrations
- `make migrate-create` : Crée une nouvelle migration
- `make migrate-status` : Affiche le statut des migrations
- `make db-reset` : Réinitialise la base de données
- `make db-tables` : Liste les tables de la base de données

## API

### Raccourcir une URL

```bash
curl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -H "X-API-Key: votre_clé_api" \
  -d '{"url": "https://example.com"}'
```

### Redirection

Accédez à `http://localhost:8080/{short_code}` pour être redirigé vers l'URL originale.

## Contribution

Les contributions sont les bienvenues ! N'hésitez pas à :
1. Forker le projet
2. Créer une branche pour votre fonctionnalité
3. Commiter vos changements
4. Pousser vers la branche
5. Ouvrir une Pull Request

## Licence

Ce projet est sous licence MIT. Voir le fichier [LICENSE](LICENSE) pour plus de détails.
