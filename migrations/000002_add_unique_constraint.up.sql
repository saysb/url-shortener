-- Ajouter une contrainte UNIQUE sur original_url
ALTER TABLE urls ADD CONSTRAINT unique_original_url UNIQUE (original_url);

-- Ajouter un index sur original_url pour améliorer les performances de recherche
CREATE INDEX idx_original_url ON urls(original_url);