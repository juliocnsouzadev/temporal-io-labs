package workflow

type WorkflowConfig struct {
	Workflow  interface{}
	TaskQueue string
	ID        string
}

func NewWorkflowConfig(w interface{}, taskQueue string, id string) *WorkflowConfig {
	return &WorkflowConfig{
		Workflow:  w,
		TaskQueue: taskQueue,
		ID:        id,
	}
}
