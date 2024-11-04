package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var (
	host     = os.Getenv("POSTGRES_HOST")
	port     = os.Getenv("POSTGRES_PORT")
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_DATABASE")

	datasource = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	campaignID = 0
	once       sync.Once
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("postgres", datasource)
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	once.Do(func() {
		_, err := db.Exec(`CREATE TABLE IF NOT EXISTS campaign(
    campaign_id SERIAL PRIMARY KEY,
    name VARCHAR (50) NOT NULL
)`)
		if err != nil {
			t.Fatalf("failed to create campaign table: %v", err)
		}

		_, err = db.Exec(`CREATE TABLE IF NOT EXISTS content(
	content_id SERIAL PRIMARY KEY,
    client_id VARCHAR (50),
	ean VARCHAR(50) NOT NULL,
	score INT NOT NULL,
	description VARCHAR(256) NOT NULL,
	url VARCHAR(256) NOT NULL,
	campaign_id INT,
	FOREIGN KEY (campaign_id) REFERENCES campaign(campaign_id)
)`)
		if err != nil {
			t.Fatalf("failed to create campaign table: %v", err)
		}
	})
	return db
}

func tearDownTestDB(db *sql.DB) {
	db.Exec("TRUNCATE TABLE content, campaign RESTART IDENTITY")
}

func TestIntegrationGetRecommendedContentByClient(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	db := setupTestDB(t)
	defer tearDownTestDB(db)

	repo := &Repo{db: db}

	_, err := db.Exec("INSERT INTO content (client_id, ean, score, description, url) VALUES ($1, $2, $3, $4, $5)",
		"clientTest", "123456", "95", "Product 1", "http://example.com/1")
	assert.NoError(t, err)

	contents, err := repo.GetRecommendedContentByClient("clientTest", 10, 0)
	assert.NoError(t, err)
	assert.NotNil(t, contents)
	assert.Len(t, contents, 1)
	assert.Equal(t, "clientTest", contents[0].ClientID)
	assert.Equal(t, "Product 1", contents[0].Description)
}

func TestIntegrationGetCampaignsContent(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	db := setupTestDB(t)
	defer tearDownTestDB(db)

	repo := &Repo{db: db}

	res := db.QueryRow("INSERT INTO campaign (name) VALUES ($1) RETURNING campaign_id;", "Test Campaign")
	err := res.Scan(&campaignID)
	assert.NoError(t, err)

	_, err = db.Exec("INSERT INTO content (campaign_id, ean, score, description, url) VALUES ($1, $2, $3, $4, $5)",
		campaignID, "654321", "90", "Campaign Product", "http://example.com/2")
	assert.NoError(t, err)

	contents, err := repo.GetCampaignsContent(10, 0)
	assert.NoError(t, err)
	assert.Len(t, contents, 1)
	assert.Equal(t, "Campaign Product", contents[0].Description)
}

func TestIntegrationCreateCampaign(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	db := setupTestDB(t)
	defer tearDownTestDB(db)

	repo := &Repo{db: db}

	campaign, err := repo.CreateCampaign("New Campaign")
	assert.NoError(t, err)
	assert.NotNil(t, campaign)
	assert.Equal(t, "New Campaign", campaign.Name)

	var name string
	err = db.QueryRow("SELECT name FROM campaign WHERE campaign_id = $1", campaign.ID).Scan(&name)
	assert.NoError(t, err)
	assert.Equal(t, "New Campaign", name)
}
