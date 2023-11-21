package workflow

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"
)

func Execute(c client.Client, config *WorkflowConfig, args ...interface{}) {
	options := client.StartWorkflowOptions{
		ID:        config.ID,
		TaskQueue: config.TaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, config.Workflow, args...)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
}
