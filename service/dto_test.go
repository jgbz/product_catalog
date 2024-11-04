package service

import (
	"testing"

	"github.com/jgbz/product_catalog/domain"
	"github.com/stretchr/testify/assert"
)

func TestGenerateIframeData(t *testing.T) {
	tests := []struct {
		name     string
		feed     *domain.ProductFeed
		expected string
	}{
		{
			name: "Empty Feed",
			feed: &domain.ProductFeed{Contents: []*domain.Content{}},
			expected: `<h2>Campaign Contents</h2><ul>
</ul>`,
		},
		{
			name: "Single Item Feed",
			feed: &domain.ProductFeed{
				Contents: []*domain.Content{
					{Description: "Product 1", URL: "http://example.com/image1.jpg"},
				},
			},
			expected: `<h2>Campaign Contents</h2><ul>
<li>
	<strong>Product 1</strong><br>
	<img src="http://example.com/image1.jpg" alt="Product Image" style="max-width: 300px; max-height: 300px; width: auto; height: auto;" />
</li>
</ul>`,
		},
		{
			name: "Multiple Items Feed",
			feed: &domain.ProductFeed{
				Contents: []*domain.Content{
					{Description: "Product 1", URL: "http://example.com/image1.jpg"},
					{Description: "Product 2", URL: "http://example.com/image2.jpg"},
				},
			},
			expected: `<h2>Campaign Contents</h2><ul>
<li>
	<strong>Product 1</strong><br>
	<img src="http://example.com/image1.jpg" alt="Product Image" style="max-width: 300px; max-height: 300px; width: auto; height: auto;" />
</li>
<li>
	<strong>Product 2</strong><br>
	<img src="http://example.com/image2.jpg" alt="Product Image" style="max-width: 300px; max-height: 300px; width: auto; height: auto;" />
</li>
</ul>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateIframeData(tt.feed)
			assert.Equal(t, tt.expected, result)
		})
	}
}
