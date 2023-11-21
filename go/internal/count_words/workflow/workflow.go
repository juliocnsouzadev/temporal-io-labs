package workflow

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"

	"github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/activity"
	"github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/tracing"
)

const (
	CountWordsTaskQueue = "count-words-tasks"
)

func CountWords(ctx workflow.Context, text string) (*activity.Reduced, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	mappedText := &activity.Mapped{}
	if val := ctx.Value(tracing.PropagateKey); val != nil {
		mappedText.TracingValues = val.(tracing.Values)
		tracing.LogDebug(ctx, mappedText.TracingValues.Data...)
	}

	err := workflow.ExecuteActivity(ctx, activity.Map, text).Get(ctx, mappedText)
	if err != nil {
		return nil, err
	}

	reducedWords := &activity.Reduced{}
	if val := ctx.Value(tracing.PropagateKey); val != nil {
		reducedWords.TracingValues = val.(tracing.Values)
		tracing.LogDebug(ctx, reducedWords.TracingValues.Data...)
	}
	err = workflow.ExecuteActivity(ctx, activity.Reduce, mappedText).Get(ctx, reducedWords)
	if err != nil {
		return nil, err
	}

	message := fmt.Sprintf("workflow completed with total of %d words mapped counted", len(mappedText.Words))
	workflow.GetLogger(ctx).Debug(message)
	return reducedWords, nil
}
