package main

import (
	"fmt"
	"log"

	"github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/workflow"
	"go.temporal.io/sdk/client"
)

var (
	lines = []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		"Aenean feugiat felis sed turpis scelerisque, at imperdiet ante viverra.",
		"Aenean nec dui nec tellus dapibus ultricies sit amet a nulla.",
		"Integer eget dolor quis dolor luctus vestibulum.",
		"Nullam et turpis ac diam pellentesque feugiat.",
		"Maecenas scelerisque lorem at diam dictum, sit amet bibendum quam sollicitudin.",
		"Sed iaculis felis vitae dui elementum rhoncus ac vitae nisi.",
		"Etiam suscipit nulla sit amet semper efficitur.",
		"Cras pulvinar dui sit amet lacus pharetra congue.",
		"Duis tristique ante a lectus venenatis, ac congue nibh euismod.",
		"Aenean accumsan nibh eu dolor gravida condimentum.",
		"Maecenas laoreet turpis in erat fermentum, nec rutrum erat facilisis.",
		"Morbi malesuada turpis sit amet fermentum volutpat.",
		"Aliquam in ligula porttitor, molestie mi sit amet, tincidunt urna.",
		"Fusce at leo sed arcu fringilla eleifend id nec libero.",
		"Proin non lectus fringilla, varius ipsum eget, vulputate dui.",
	}
)

func main() {

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	for i, line := range lines {

		id := fmt.Sprintf("count-words-%d", i)
		cfg := workflow.NewWorkflowConfig(workflow.CountWords, workflow.CountWordsTaskQueue, id)
		workflow.Execute(c, cfg, line)
	}

}