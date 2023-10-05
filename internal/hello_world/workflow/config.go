package workflow

type WorkflowConfig struct {
	Workflow  interface{}
	TaskQueue string
}

func NewWorkflowConfig(w interface{}, taskQueue string) *WorkflowConfig {
	return &WorkflowConfig{
		Workflow:  w,
		TaskQueue: taskQueue,
	}
}
