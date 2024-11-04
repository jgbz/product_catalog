BEGIN;

CREATE TABLE IF NOT EXISTS content(
	content_id SERIAL PRIMARY KEY,
    client_id VARCHAR (50),
	ean VARCHAR(50) NOT NULL,
	score INT NOT NULL,
	description VARCHAR(256) NOT NULL,
	url VARCHAR(256) NOT NULL,
	campaign_id INT,
	FOREIGN KEY (campaign_id) REFERENCES campaign(campaign_id)
);

COMMIT;