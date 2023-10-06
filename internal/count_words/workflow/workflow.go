package workflow

import (
	"fmt"
	"time"

	"github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/activity"
	"go.temporal.io/sdk/workflow"
)

const (
	CountWordsTaskQueue = "count-words-tasks"
)

func CountWords(ctx workflow.Context, text string) (map[string]int, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var mappedText activity.Mapped
	err := workflow.ExecuteActivity(ctx, activity.Map, text).Get(ctx, &mappedText)
	if err != nil {
		return nil, err
	}

	var reducedWords activity.Reduced
	err = workflow.ExecuteActivity(ctx, activity.Reduce, mappedText).Get(ctx, &reducedWords)
	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf("Total of %d words mapped from %s", len(mappedText.Words), text)
	workflow.GetLogger(ctx).Info(message)
	return reducedWords.WordCount, nil
}
