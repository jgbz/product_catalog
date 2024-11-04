package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/jgbz/product_catalog/domain"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// Validate if FeedService implements the Service interface
var _ Repository = (*Repo)(nil)

type Repo struct {
	db *sql.DB
}

// NewRepository connects to a postgres database and try to ping it, returns a wrapper.
func NewRepository() (*Repo, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE")

	datasource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", datasource)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Repo{
		db: db,
	}, nil
}

// RunMigration executes all the migrations file in the database.
func (repo *Repo) RunMigration() error {
	driver, err := postgres.WithInstance(repo.db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	// If its get a error other than ErrNoChange, returns it.
	if err := m.Up(); err != migrate.ErrNoChange {
		return err
	}
	return nil
}

// GetRecommendedContentByClient retrieves recommended content that belongs to the specified clientID, using limit and offset for pagination.
func (repo *Repo) GetRecommendedContentByClient(clientID string, limit, offset int) ([]*domain.Content, error) {
	query := `
	SELECT client_id, ean, score, description, url
	FROM content
	WHERE client_id = $1 
	ORDER BY score DESC
	LIMIT $2 OFFSET $3`

	stmt, err := repo.db.PrepareContext(context.Background(), query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(clientID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	contents := make([]*domain.Content, 0)
	for rows.Next() {
		content := &domain.Content{}
		if err := rows.Scan(&content.ClientID, &content.EAN, &content.Score, &content.Description, &content.URL); err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}

	return contents, nil
}

// GetCampaignsContent retrieves all the content that belongs to campaigns, using limit and offset for pagination.
func (repo *Repo) GetCampaignsContent(limit, offset int) ([]*domain.Content, error) {
	query := `
	SELECT campaign_id, ean, score, description, url
	FROM content
	WHERE campaign_id IN (select campaign_id from campaign)
	ORDER BY score DESC
	LIMIT $1 OFFSET $2`

	stmt, err := repo.db.PrepareContext(context.Background(), query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	contents := make([]*domain.Content, 0)
	for rows.Next() {
		content := &domain.Content{}
		if err := rows.Scan(&content.CampaignID, &content.EAN, &content.Score, &content.Description, &content.URL); err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}

	return contents, nil
}

// CreateCampaign creates a new campaign record.
func (repo *Repo) CreateCampaign(name string) (*domain.Campaign, error) {
	query := "INSERT INTO campaign (name) VALUES ($1) RETURNING campaign_id"
	stmt, err := repo.db.PrepareContext(context.Background(), query)
	if err != nil {
		return nil, err
	}

	lastInsertId := 0
	err = stmt.QueryRowContext(context.Background(), name).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	return &domain.Campaign{ID: strconv.Itoa(lastInsertId), Name: name}, nil
}

// Health pings the database to check if the connection is alive.
func (repo *Repo) Health() error {
	return repo.db.Ping()
}
