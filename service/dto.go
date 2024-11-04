package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jgbz/product_catalog/domain"
)

var (
	ErrFailedToParseFeed = errors.New("failed to parse feed content")
)

// CampaignRequest is the expected structure for creating a campaign.
type CampaignRequest struct {
	Name string `json:"name"`
}

// CampaignResponse is the expected response structure when creating a campaign.
type CampaignResponse struct {
	CampaignID string `json:"campaign_id"`
}

// ToCampaignResponse parses the domain.Campaign into a CampaignResponse.
func ToCampaignResponse(campaign *domain.Campaign) CampaignResponse {
	return CampaignResponse{
		CampaignID: campaign.ID,
	}
}

// FeedRequest is the expected structure then retrieving a feed.
type FeedRequest struct {
	ClientID string `json:"client_id,omitempty"`
	Page     int
}

// FeedResponse is the expected response structure in a feed retrieval.
type FeedResponse struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
	Data       string `json:"data,omitempty"`
}

// ToFeedResponse parses the operation code, message and the domain.Feed into the FeedResponse structure.
func ToFeedResponse(code int, message string, feed *domain.ProductFeed) (FeedResponse, error) {
	return FeedResponse{
		StatusCode: code,
		Message:    message,
		Data:       generateIframeData(feed),
	}, nil
}

// generateIframeData builds the html string of the iframe, iterating over every content on the feed and generating a list entry for each one.
func generateIframeData(feed *domain.ProductFeed) string {
	var sb strings.Builder

	var itemHTML = `<li>
	<strong>%s</strong><br>
	<img src="%s" alt="Product Image" style="max-width: 300px; max-height: 300px; width: auto; height: auto;" />
</li>
`

	sb.WriteString("<h2>Campaign Contents</h2><ul>\n")
	for _, c := range feed.Contents {
		// not real production code, but its only to see that the campaigns are in order.
		if c.CampaignID != "" {
			c.Description = fmt.Sprintf("%v#%s", c.CampaignID, c.Description)
		}
		sb.WriteString(fmt.Sprintf(itemHTML, c.Description, c.URL))
	}
	sb.WriteString("</ul>")

	return sb.String()
}
