package main

import (
	"log"

	"github.com/juliocnsouzadev/temporal-io-labs/internal/hello_world/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	cfg := workflow.NewWorkflowConfig(workflow.HelloWorldWorkflow, "hello-world")

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, cfg.TaskQueue, worker.Options{})
	w.RegisterWorkflow(cfg.Workflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
