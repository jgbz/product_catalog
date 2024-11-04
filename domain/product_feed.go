package domain

// ProductFeed represents the client feed.
type ProductFeed struct {
	ClientID string     `json:"clientId"`
	Contents []*Content `json:"contents"`
}

// Content represents a product that can be from a campaign, or from a clients personal feed.
type Content struct {
	ClientID    string `json:"client_id,omitempty"`
	EAN         string `json:"ean"`
	Score       string `json:"score"`
	Description string `json:"description"`
	URL         string `json:"url"`
	CampaignID  string `json:"campaign_id,omitempty"`
}

// Campaign represents a sponsored set of contents.
type Campaign struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
