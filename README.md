# Temporal Workflos with Go
> ### [Temporal](https://temporal.io/)

## What is Temporal?
In short, Temporal is a platform that guarantees the durable execution of your application code. It allows you to develop as if failures don't even exist. Your application will run reliably even if it encounters problems, such as network outages or server crashes, which would be catastrophic for a typical application. The Temporal platform handles these types of problems, allowing you to focus on the business logic, instead of writing application code to detect and recover from failures.

## Workflows
Temporal applications are built using an abstraction called Workflows. You'll develop those Workflows by writing code in a general-purpose programming language such as Go, Java, TypeScript, or Python. The code you write is the same code that will be executed at runtime, so you can use your favorite tools and libraries to develop Temporal Workflows.

Temporal Workflows are resilient. They can run—and keeping running—for years, even if the underlying infrastructure fails. If the application itself crashes, Temporal will automatically recreate its pre-failure state so it can continue right where it left off.

## Install the SDK
```bash
go get go.temporal.io/sdk
```

## CLI
Temporal provides a command-line interface (CLI), `tctl`, which allows you to interact with a cluster. 

### Installing 
```bash
brew install tctl
```

### Checking available commands
```bash
tctl workflow --help
```

### Running Workflow Worker
```bash
go run ./cmd/hello_world/main.go 
```

### Staring Workflow
```bash
tctl workflow start \
    --workflow_type HelloWorldWorkflow \
    --taskqueue hello-world \
    --workflow_id hello-world-workflow \
    --input '"Julio"'
```

## Workflow History
- UI
- Using `tctl`:
```bash
tctl wf show --workflow_id my-first-workflow
```

Example Output:
```bash
  1  WorkflowExecutionStarted    {WorkflowType:{Name:GreetSomeone},
                                  ParentInitiatedEventId:0, TaskQueue:{Name:greeting-tasks,
                                  Kind:Normal}, Input:["Donna"],
                                  WorkflowExecutionTimeout:0s, WorkflowRunTimeout:0s,
                                  WorkflowTaskTimeout:10s, Initiator:Unspecified,
                                  OriginalExecutionRunId:e8f9217e-344e-4f7b-98bc-7703bc8c7c76,
                                  Identity:tctl@twwmbp,
                                  FirstExecutionRunId:e8f9217e-344e-4f7b-98bc-7703bc8c7c76,
                                  Attempt:1, FirstWorkflowTaskBackoff:0s,
                                  ParentInitiatedEventVersion:0}
  2  WorkflowTaskScheduled       {TaskQueue:{Name:greeting-tasks,
                                  Kind:Normal},
                                  StartToCloseTimeout:10s,
                                  Attempt:1}
  3  WorkflowTaskStarted         {ScheduledEventId:2, Identity:93592@twwmbp@,
                                  RequestId:10535889-9c10-4073-b38f-4876bbae4db3,
                                  SuggestContinueAsNew:false, HistorySizeBytes:0}
  4  WorkflowTaskCompleted       {ScheduledEventId:2, StartedEventId:3,
                                  Identity:93592@twwmbp@,
                                  BinaryChecksum:202d5177234b6ec7b33e3de1b92f2f5f}
  5  WorkflowExecutionCompleted  {Result:["Hello Donna!"],
                                  WorkflowTaskCompletedEventId:4}
```

## Activities

In Temporal, you can use Activities to encapsulate business logic that is prone to failure. Unlike the Workflow Definition, there is no requirement for an Activity Definition to be deterministic.

In general, any operation that introduces the possibility of failure should be done as part of an Activity, rather than as part of the Workflow directly. While Activities are executed as part of Workflow Execution, they have an important characteristic: they're retried if they fail. If you have an extensive Workflow that needs to access a service, and that service happens to become unavailable, you don't want to re-run the entire Workflow. Instead, you just want to retry the part that failed, so you can define that code in an Activity and reference it in your Workflow Definition.

### Example with Activity

#### Run Worker
```bash
go run ./cmd/count_works/woker/main.go
```

#### Run Execution Example
```bash
go run ./cmd/count_works/example/main.go
```

or

#### Run Uploading a Text File via API
```bash
go run ./cmd/count_works/example/main.go
```

**Upload the file with CURL**
```bash
curl -X POST -F "file=@path_to_file" http://localhost:8080/upload
```

**Upload the file with Postman**

1. **Open Postman**: Launch the Postman application on your computer.
2. **Create a New Request**: Click the `+` button or "New" to create a new request.
3. **Set the Request Type**: In the dropdown next to the URL bar (which usually defaults to GET), select `POST`.
4. **Enter the URL**: Input the URL where you want to upload the file, e.g., `http://localhost:9090/upload`.
5. **Select `Body` Tab**: Below the URL bar, you'll find several tabs such as Params, Authorization, Headers, etc. Select the `Body` tab.
6. **Choose `form-data`**: In the `Body` tab, you'll see options like `form-data`, `x-www-form-urlencoded`, `raw`, `binary`, etc. Select `form-data`.
7. **Add File**: 
   - On the left-hand side, you'll be able to input the key (name of the form field). Enter `file` (or whatever the server expects as the field name).
   - On the right-hand side, instead of choosing `Text`, click on the dropdown and choose `File`.
   - An input field labeled "Select Files" will appear on the right. Click on it and choose the file you want to upload from your computer.
8. **Send the Request**: Click the blue `Send` button. Postman will make a POST request with the file attached.
9. **Inspect the Response**: After sending, Postman will display the server's response in the lower half of the window. This can help you understand if the upload was successful or if there were any issues.
