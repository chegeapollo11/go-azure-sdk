
## `github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/storageaccounts` Documentation

The `storageaccounts` SDK allows for interaction with the Azure Resource Manager Service `databoxedge` (API Version `2020-12-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/storageaccounts"
```


### Client Initialization

```go
client := storageaccounts.NewStorageAccountsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `StorageAccountsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "storageAccountValue")

payload := storageaccounts.StorageAccount{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `StorageAccountsClient.Delete`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "storageAccountValue")

if err := client.DeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `StorageAccountsClient.Get`

```go
ctx := context.TODO()
id := storageaccounts.NewStorageAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue", "storageAccountValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `StorageAccountsClient.ListByDataBoxEdgeDevice`

```go
ctx := context.TODO()
id := storageaccounts.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue")

// alternatively `client.ListByDataBoxEdgeDevice(ctx, id)` can be used to do batched pagination
items, err := client.ListByDataBoxEdgeDeviceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
