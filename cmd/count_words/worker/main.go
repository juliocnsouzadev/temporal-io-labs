package main

import (
	"log"

	"github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/activity"
	"github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/workflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, workflow.CountWordsTaskQueue, worker.Options{})

	w.RegisterWorkflow(workflow.CountWords)
	w.RegisterActivity(activity.Map)
	w.RegisterActivity(activity.Reduce)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
