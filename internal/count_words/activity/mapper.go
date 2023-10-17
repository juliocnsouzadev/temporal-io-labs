package activity

import (
	"context"
	"github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/tracing"
	"strings"
)

type Mapped struct {
	Words         []string       `json:"words"`
	TracingValues tracing.Values `json:"tracingValues"`
}

func Map(ctx context.Context, text string) (*Mapped, error) {
	text = strings.ToLower(text)
	words := strings.Split(text, " ")

	result := &Mapped{
		Words: words,
	}
	if val := ctx.Value(tracing.PropagateKey); val != nil {
		result.TracingValues = val.(tracing.Values)
	}
	return result, nil
}
