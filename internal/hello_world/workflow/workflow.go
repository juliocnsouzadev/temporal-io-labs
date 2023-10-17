package workflow

import (
	"fmt"

	"go.temporal.io/sdk/workflow"
)

func HelloWorldWorkflow(ctx workflow.Context, params string) (string, error) {
	message := fmt.Sprintf("Hello World! => %s", params)
	workflow.GetLogger(ctx).Debug(message)
	return message, nil
}
