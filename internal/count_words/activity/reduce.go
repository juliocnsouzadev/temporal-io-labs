package activity

import (
	"context"
)

type Reduced struct {
	WordCount map[string]int `json:"word_count"`
}

func Reduce(ctx context.Context, data *Mapped) (*Reduced, error) {
	wordCount := make(map[string]int)
	for _, word := range data.Words {
		wordCount[word]++
	}

	return &Reduced{
		WordCount: wordCount,
	}, nil
}
