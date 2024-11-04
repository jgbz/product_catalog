package database

import "github.com/jgbz/product_catalog/domain"

// Repository interface is the abstraction layer for database operations.
//
//go:generate mockgen -destination mock/database.go github.com/jgbz/product_catalog/database Repository
type Repository interface {
	CreateCampaign(name string) (*domain.Campaign, error)
	GetCampaignsContent(limit, offset int) ([]*domain.Content, error)
	GetRecommendedContentByClient(clientID string, limit, offset int) ([]*domain.Content, error)
	Health() error
}
