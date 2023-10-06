package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"

	"github.com/juliocnsouzadev/temporal-io-labs/internal/count_words/workflow"
	"go.temporal.io/sdk/client"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data to retrieve the file
	err := r.ParseMultipartForm(2 << 20) // 10 MB limit
	if err != nil {
		msg := fmt.Sprintf("Unable to parse form: %s", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	// Get the uploaded file
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		id := fmt.Sprintf("count-words-%d", i)
		cfg := workflow.NewWorkflowConfig(workflow.CountWords, workflow.CountWordsTaskQueue, id)
		workflow.Execute(c, cfg, scanner.Text())
		i++
	}

	if err := scanner.Err(); err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("File uploaded successfully!"))
}

func main() {
	http.HandleFunc("/upload", uploadHandler)

	fmt.Println("Server started on :9090")
	http.ListenAndServe(":9090", nil)
}
