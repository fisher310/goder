package model

import "github.com/fisher310/goder/crawler/engine"

type SearchResult struct {
	Hits  int
	Start int
	Items []engine.Item
}
