package tracing

import (
	"fmt"
	"go.temporal.io/sdk/workflow"
)

func LogDebug(ctx workflow.Context, data ...KeyValue) {
	msg := "[ "
	for _, d := range data {
		msg += fmt.Sprintf("%v,", d)
	}
	msg += " ]"
	workflow.GetLogger(ctx).Debug(msg)
}
