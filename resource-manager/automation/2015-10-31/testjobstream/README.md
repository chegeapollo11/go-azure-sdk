
## `github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/testjobstream` Documentation

The `testjobstream` SDK allows for interaction with the Azure Resource Manager Service `automation` (API Version `2015-10-31`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/automation/2015-10-31/testjobstream"
```


### Client Initialization

```go
client := testjobstream.NewTestJobStreamClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TestJobStreamClient.Get`

```go
ctx := context.TODO()
id := testjobstream.NewStreamID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountValue", "jobIdValue", "jobStreamIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TestJobStreamClient.ListByTestJob`

```go
ctx := context.TODO()
id := testjobstream.NewRunbookID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountValue", "runbookValue")

// alternatively `client.ListByTestJob(ctx, id, testjobstream.DefaultListByTestJobOperationOptions())` can be used to do batched pagination
items, err := client.ListByTestJobComplete(ctx, id, testjobstream.DefaultListByTestJobOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
