package fetcher

import (
	"testing"
)

func TestFetcher(t *testing.T) {
	tests := []struct {
		url     string
		bodyLen int
	}{
		{"http://product.dangdang.com/22610008.html", 100000},
		{"http://product.dangdang.com/23578344.html", 100000},
		{"http://product.dangdang.com/22911745.html", 100000},
	}
	for k, test := range tests {
		body, err := Fetcher(test.url)
		if err != nil {
			t.Errorf("Number %d test error.Error of fetcher: %v", k, err)
		}
		if len(body) < test.bodyLen {
			t.Errorf("Number %d test error.Expected body length should over 10000, but got: %d", k, len(body))
		}
	}
}
