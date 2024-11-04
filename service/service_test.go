package service

import (
	"testing"

	mock_database "github.com/jgbz/product_catalog/database/mock"
	"github.com/jgbz/product_catalog/domain"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetFeedByClient_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_database.NewMockRepository(ctrl)
	feedService := &FeedService{repo: mockRepo}

	req := &FeedRequest{ClientID: "123", Page: 1}

	expectedContent := []*domain.Content{{ClientID: "123", Description: "test", URL: "url"}}
	mockRepo.EXPECT().GetRecommendedContentByClient("123", limit, 0).Return(expectedContent, nil)
	mockRepo.EXPECT().GetCampaignsContent(limit, 0).Return([]*domain.Content{}, nil)

	expectedFeedResponse, err := ToFeedResponse(StatusOK, "message", &domain.ProductFeed{ClientID: "123", Contents: expectedContent})
	assert.NoError(t, err)

	resp, err := feedService.GetFeedByClient(req)

	assert.NoError(t, err)
	assert.Equal(t, StatusOK, resp.StatusCode)
	assert.Equal(t, expectedFeedResponse, resp)
}

func TestGetFeedByClient_FailedToGetRecommendendContentByClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_database.NewMockRepository(ctrl)
	feedService := &FeedService{repo: mockRepo}

	req := &FeedRequest{ClientID: "123", Page: 1}

	mockRepo.EXPECT().GetRecommendedContentByClient("123", limit, 0).Return(nil, ErrFailedToRetrieveRecommendedContents)
	// mockRepo.EXPECT().GetCampaignsContent(limit, 0).Return([]*domain.Content{}, nil)

	expectedFeedResponse := FeedResponse{StatusCode: StatusInternalServerError}

	resp, err := feedService.GetFeedByClient(req)

	assert.Error(t, err)
	assert.Equal(t, StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, expectedFeedResponse, resp)
}

func TestGetFeedByClient_FailedToGetCampaignsContent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_database.NewMockRepository(ctrl)
	feedService := &FeedService{repo: mockRepo}

	req := &FeedRequest{ClientID: "123", Page: 1}

	mockRepo.EXPECT().GetRecommendedContentByClient("123", limit, 0).Return([]*domain.Content{}, nil)
	mockRepo.EXPECT().GetCampaignsContent(limit, 0).Return(nil, ErrFailedToRetrieveCampaignContents)

	expectedFeedResponse := FeedResponse{StatusCode: StatusInternalServerError}

	resp, err := feedService.GetFeedByClient(req)

	assert.Error(t, err)
	assert.Equal(t, StatusInternalServerError, resp.StatusCode)
	assert.Equal(t, expectedFeedResponse, resp)
}

func TestCreateCampaign_EmptyName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_database.NewMockRepository(ctrl)
	feedService := &FeedService{repo: mockRepo}

	req := &CampaignRequest{Name: ""}

	resp, err := feedService.CreateCampaign(req)

	assert.Error(t, err)
	assert.Equal(t, ErrEmptyCampaignName, err)
	assert.Empty(t, resp)
}

func TestCreateCampaign_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_database.NewMockRepository(ctrl)
	feedService := &FeedService{repo: mockRepo}

	req := &CampaignRequest{Name: "New Campaign"}
	expectedCampaign := &domain.Campaign{ID: "1", Name: "New Campaign"}

	mockRepo.EXPECT().CreateCampaign(req.Name).Return(expectedCampaign, nil)

	resp, err := feedService.CreateCampaign(req)

	assert.NoError(t, err)
	assert.Equal(t, expectedCampaign.ID, resp.CampaignID)
}

func TestRepoHealth_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_database.NewMockRepository(ctrl)
	feedService := &FeedService{repo: mockRepo}

	mockRepo.EXPECT().Health().Return(nil)

	err := feedService.RepoHealth()

	assert.NoError(t, err)
}
