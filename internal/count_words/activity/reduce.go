package activity

import (
	"context"
	"github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/tracing"
	"log"
)

type Reduced struct {
	WordCount     map[string]int `json:"word_count"`
	TracingValues tracing.Values `json:"tracingValues"`
}

func Reduce(ctx context.Context, data *Mapped) (*Reduced, error) {
	wordCount := make(map[string]int)
	for _, word := range data.Words {
		wordCount[word]++
	}

	result := Reduced{
		WordCount: wordCount,
	}
	if val := ctx.Value(tracing.PropagateKey); val != nil {
		if tracingValues, ok := val.(tracing.Values); ok {
			result.TracingValues = tracingValues
		} else {
			log.Println("no propagate key found in context [reducing]")
		}
	}
	return &result, nil
}
