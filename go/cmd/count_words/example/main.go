package main

import (
	"context"
	"fmt"
	"log"
	"slices"
	"strconv"
	"time"

	"github.com/google/uuid"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/contrib/opentracing"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/interceptor"

	"go.temporal.io/sdk/client"
	temporalWorkflow "go.temporal.io/sdk/workflow"

	"github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/activity"
	tracing2 "github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/tracing"
	my_workflow "github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/workflow"
)

var (
	lines = []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Aenean feugiat felis sed turpis scelerisque, at imperdiet ante viverra.",
		"Aenean nec dui nec tellus dapibus ultricies sit amet a nulla.",
		"Integer eget dolor quis dolor luctus vestibulum.",
		"Nullam et turpis ac diam pellentesque feugiat.",
		"Maecenas scelerisque lorem at diam dictum, sit amet bibendum quam sollicitudin.",
		"Sed iaculis felis vitae dui elementum rhoncus ac vitae nisi.",
		"Etiam suscipit nulla sit amet semper efficitur.",
		"Cras pulvinar dui sit amet lacus pharetra congue.",
		"Duis tristique ante a lectus venenatis, ac congue nibh euismod.",
		"Aenean accumsan nibh eu dolor gravida condimentum.",
		"Maecenas laoreet turpis in erat fermentum, nec rutrum erat facilisis.",
		"Morbi malesuada turpis sit amet fermentum volutpat.",
		"Aliquam in ligula porttitor, molestie mi sit amet, tincidunt urna.",
		"Fusce at leo sed arcu fringilla eleifend id nec libero.",
		"Proin non lectus fringilla, varius ipsum eget, vulputate dui.",
	}
)

func main() {
	// Set tracer which will be returned by opentracing.GlobalTracer().
	closer, err := tracing2.SetJaegerGlobalTracer("word-count")
	if err != nil {
		log.Fatalf("Failed creating tracer: %v", err)
	}
	defer func() {
		if err := closer.Close(); err != nil {
			log.Fatalf("Failed to close tracer: %v", err)
		}
	}()

	// Create interceptor
	tracingInterceptor, err := opentracing.NewInterceptor(opentracing.TracerOptions{})
	if err != nil {
		log.Fatalf("Failed creating interceptor: %v", err)
	}

	c, err := client.Dial(client.Options{
		HostPort:           client.DefaultHostPort,
		Interceptors:       []interceptor.ClientInterceptor{tracingInterceptor},
		ContextPropagators: []temporalWorkflow.ContextPropagator{tracing2.NewContextPropagator()},
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	cId, _ := uuid.NewUUID()
	correlationId := my_workflow.WorkflowMetadata{
		Key:   "correlationId",
		Value: cId.String(),
	}

	for _, line := range lines {
		// execute the first workflow
		workflowID, ctx := runWorkflow("cw01", line, correlationId, c)

		// get the result of the first workflow
		mappedText := getMapResult(c, ctx, workflowID)

		// execute the second workflow
		newLine := reversedText(mappedText)
		workflowID, ctx = runWorkflow("cw02", newLine, correlationId, c)
		break // just one iteration for the example
	}

}

func reversedText(mappedText *activity.Mapped) string {
	newLine := ""
	slices.Reverse(mappedText.Words)
	for _, word := range mappedText.Words {
		newLine += " " + word
	}
	return newLine
}

func getMapResult(c client.Client, ctx context.Context, workflowID string) *activity.Mapped {
	// Get the workflow execution history
	wfHistory := c.GetWorkflowHistory(
		ctx,
		workflowID,
		"",
		false,
		enums.HISTORY_EVENT_FILTER_TYPE_ALL_EVENT)

	mappedText := &activity.Mapped{}
	for wfHistory.HasNext() {
		event, err := wfHistory.Next()
		if err != nil {
			log.Fatalf("Failed to read history: %v", err)
		}
		eventType := event.GetEventType()
		log.Printf("Event type: %v", eventType)

		if eventType == enums.EVENT_TYPE_ACTIVITY_TASK_COMPLETED {
			log.Printf("Activity task completed: %v", event.EventId)
			attributes := event.GetActivityTaskCompletedEventAttributes()

			dataConverter := converter.GetDefaultDataConverter()
			err := dataConverter.FromPayloads(attributes.GetResult(), mappedText)
			if err != nil || mappedText.Words == nil {
				log.Printf("Failed to decode result attributes.GetResult(): %v", err)
				continue // skip to the next event since it is not activity.Mapped object
			}
			log.Printf("Mapped words: %v", mappedText.Words)
			break
		}
	}
	return mappedText
}

func runWorkflow(prefix, line string, correlationId my_workflow.WorkflowMetadata, c client.Client) (string, context.Context) {
	milli := time.Now().UnixMilli()
	workflowID := fmt.Sprintf("%s-%d", prefix, milli)
	textSize := my_workflow.WorkflowMetadata{
		Key:   "textSize",
		Value: strconv.Itoa(len(line)),
	}
	cfg := my_workflow.NewWorkflowConfig(my_workflow.CountWords, my_workflow.CountWordsTaskQueue, workflowID, correlationId, textSize)
	ctx := my_workflow.Execute(c, cfg, line)

	return workflowID, ctx
}
