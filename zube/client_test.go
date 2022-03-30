package zube

import (
	"testing"
)

func TestQueryEncode(t *testing.T) {
	var examples = []Query{
		{Pagination: Pagination{Page: "1"}, Order: Order{By: "id"}},
		{Order: Order{By: "title", Direction: "asc"}},
		{Filter: Filter{Where: map[string]any{"id": 123}, Select: []string{"id", "title"}}},
	}

	var wants = []string{
		"order%5Bby%5D=id&order%5Bdirection%5D=&page=1&per_page=",
		"order%5Bby%5D=title&order%5Bdirection%5D=asc&page=&per_page=",
		"order%5Bby%5D=&order%5Bdirection%5D=&page=&per_page=&select%5B%5D=id&select%5B%5D=title&where%5Bid%5D=123"
	}

	for i := 0; i < 3; i++ {
		res := examples[i].Encode()
		if res != wants[i] {
			t.Errorf("got %s expected: %s", res, wants[i])
		}
	}
}
