package tracing

import (
	"fmt"
	"go.temporal.io/sdk/workflow"
	"strings"
)

func LogDebug(ctx workflow.Context, data ...KeyValue) {
	var msg strings.Builder
	for _, d := range data {
		msg.WriteString(fmt.Sprintf("%v ", d))
	}
	workflow.GetLogger(ctx).Debug(msg.String())
}
