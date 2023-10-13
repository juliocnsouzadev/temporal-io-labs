package workflow

type WorkflowConfig struct {
	Workflow  interface{}
	TaskQueue string
	ID        string
	Metadata  map[string]interface{}
}

type WorkflowMetadata struct {
	Key   string
	Value interface{}
}

func NewWorkflowConfig(w interface{}, taskQueue string, id string, metadata ...WorkflowMetadata) *WorkflowConfig {
	meta := make(map[string]interface{}, len(metadata))
	for _, data := range metadata {
		meta[data.Key] = data.Value
	}
	return &WorkflowConfig{
		Workflow:  w,
		TaskQueue: taskQueue,
		ID:        id,
		Metadata:  meta,
	}
}
