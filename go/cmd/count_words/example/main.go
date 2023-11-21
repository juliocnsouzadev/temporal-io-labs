package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"go.temporal.io/sdk/contrib/opentracing"
	"go.temporal.io/sdk/interceptor"

	"go.temporal.io/sdk/client"
	temporalWorkflow "go.temporal.io/sdk/workflow"

	tracing2 "github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/tracing"
	workflow2 "github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/workflow"
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
	correlationId := workflow2.WorkflowMetadata{
		Key:   "correlationId",
		Value: cId.String(),
	}

	for _, line := range lines {
		milli := time.Now().UnixMilli()
		id := fmt.Sprintf("cw-%d", milli)
		textSize := workflow2.WorkflowMetadata{
			Key:   "textSize",
			Value: strconv.Itoa(len(line)),
		}
		cfg := workflow2.NewWorkflowConfig(workflow2.CountWords, workflow2.CountWordsTaskQueue, id, correlationId, textSize)
		workflow2.Execute(c, cfg, line)
	}

}
