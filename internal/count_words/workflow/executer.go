package workflow

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func Execute(c client.Client, config *WorkflowConfig, args ...interface{}) {

	options := buildOptions(config)

	we, err := c.ExecuteWorkflow(context.Background(), options, config.Workflow, args...)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	logWorkflowStart(we, options)

	var result map[string]int
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
}

func buildOptions(config *WorkflowConfig) client.StartWorkflowOptions {

	options := client.StartWorkflowOptions{}

	if config != nil {
		options.ID = config.ID
		options.TaskQueue = config.TaskQueue

		if config.Metadata != nil {
			options.SearchAttributes = config.Metadata
		}
	}
	return options
}

func logWorkflowStart(we client.WorkflowRun, options client.StartWorkflowOptions) {
	workflowId := we.GetID()
	runId := we.GetRunID()

	var correlationId interface{}
	if options.SearchAttributes != nil {
		var ok bool
		if correlationId, ok = options.SearchAttributes["correlationId"]; !ok {
			correlationId = "none"
		}
	} else {
		correlationId = "none"
	}

	log.Println("Started workflow", "WorkflowID", workflowId, "RunID", runId, "CorrelationId", correlationId)
}
