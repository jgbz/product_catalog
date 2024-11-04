package service

import (
	"errors"

	"github.com/jgbz/product_catalog/database"
	"github.com/jgbz/product_catalog/domain"
)

// Validate if FeedService implements the Service interface
var _ Service = (*FeedService)(nil)

var (
	ErrEmptyCampaignName                   = errors.New("campaign name can't be blank")
	ErrFailedToRetrieveRecommendedContents = errors.New("unable to retrieve recommended content")
	ErrFailedToRetrieveCampaignContents    = errors.New("unable to retrieve campaign content")
)

const (
	// Arbitrary value for pagination
	limit = 5

	StatusOK = iota + 10000
	StatusInternalServerError
	StatusClientError
)

// Service interface is the abstraction layer for operations related to the client feed.
type Service interface {
	GetFeedByClient(req *FeedRequest) (FeedResponse, error)
	CreateCampaign(req *CampaignRequest) (CampaignResponse, error)
	RepoHealth() error
}

// FeedService implements the Service interface and wraps the repository.
type FeedService struct {
	repo database.Repository
}

// NewFeedService returns a instance of the FeedService
func NewFeedService(repo database.Repository) Service {
	return &FeedService{
		repo: repo,
	}
}

// GetFeedByClient receives a request containing the clients id and the current page
// It retrieves all the recomended client content, and all the campaign contents
// Then merges then into a single array, based on the content.Score and it is ordered by interspersing recommended content and campaign content.
func (fs *FeedService) GetFeedByClient(req *FeedRequest) (FeedResponse, error) {
	// Define pagination constants
	offset := (req.Page - 1) * limit // Calculate offset

	recommendedContent, err := fs.repo.GetRecommendedContentByClient(req.ClientID, limit, offset)
	if err != nil {
		return FeedResponse{StatusCode: StatusInternalServerError}, errors.Join(ErrFailedToRetrieveRecommendedContents, err)
	}
	campaignContent, err := fs.repo.GetCampaignsContent(limit, offset)
	if err != nil {
		return FeedResponse{StatusCode: StatusInternalServerError}, errors.Join(ErrFailedToRetrieveCampaignContents, err)
	}

	feed := &domain.ProductFeed{}
	feed.Contents = mergeContents(recommendedContent, campaignContent)

	return ToFeedResponse(StatusOK, "message", feed)
}

// CreateCampaign receives a CampaignRequest containing the campaign name
// Returns a CampaignReponse with the ID of the newly created campaign.
func (fs *FeedService) CreateCampaign(req *CampaignRequest) (CampaignResponse, error) {
	if req.Name == "" {
		return CampaignResponse{}, ErrEmptyCampaignName
	}

	campaign, err := fs.repo.CreateCampaign(req.Name)
	if err != nil {
		return CampaignResponse{}, err
	}

	return ToCampaignResponse(campaign), nil
}

func (fs *FeedService) RepoHealth() error {
	return fs.repo.Health()
}

// mergeContents receive two arrays, one of recommended contents, and other or campaign contents.
// Then orders alternating between then two.
func mergeContents(recomended, campaigns []*domain.Content) []*domain.Content {
	// Determine the length of the new slice
	maxLen := len(recomended)
	if len(campaigns) > maxLen {
		maxLen = len(campaigns)
	}

	merged := make([]*domain.Content, 0, len(recomended)+len(campaigns))

	for i := 0; i < maxLen; i++ {
		if i < len(recomended) {
			merged = append(merged, recomended[i])
		}
		if i < len(campaigns) {
			merged = append(merged, campaigns[i])
		}
	}

	return merged
}
