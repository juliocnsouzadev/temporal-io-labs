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

### Checking avaliable commands
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