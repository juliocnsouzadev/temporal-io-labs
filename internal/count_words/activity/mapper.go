package activity

import (
	"context"
	"github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/tracing"
	"log"
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
		if tracingValues, ok := val.(tracing.Values); ok {
			result.TracingValues = tracingValues
		} else {
			log.Println("no propagate key found in context [mapping]")
		}
	}
	return result, nil
}
