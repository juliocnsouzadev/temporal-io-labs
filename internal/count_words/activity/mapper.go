package activity

import (
	"context"
	"strings"
)

type Mapped struct {
	Words []string `json:"words"`
}

func Map(ctx context.Context, text string) (*Mapped, error) {
	text = strings.ToLower(text)
	words := strings.Split(text, " ")

	return &Mapped{
		Words: words,
	}, nil
}
