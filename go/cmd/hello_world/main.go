package main

import (
	"fmt"
	"log"

	"go.temporal.io/sdk/client"

	"github.com/juliocnsouzadev/temporal-io-labs/internal/hello_world/workflow"
)

func main() {

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	for i := 0; i < 10; i++ {
		id := fmt.Sprintf("hello-world-%d", i)
		cfg := workflow.NewWorkflowConfig(workflow.HelloWorldWorkflow, "hello-world", id)
		txt := fmt.Sprintf("Hello World %d", i)
		workflow.Execute(c, cfg, txt)
	}
}
