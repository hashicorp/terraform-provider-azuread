
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/applicationtemplates/stable/applicationtemplate` Documentation

The `applicationtemplate` SDK allows for interaction with Microsoft Graph `applicationtemplates` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/applicationtemplates/stable/applicationtemplate"
```


### Client Initialization

```go
client := applicationtemplate.NewApplicationTemplateClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationTemplateClient.GetApplicationTemplate`

```go
ctx := context.TODO()
id := applicationtemplate.NewApplicationTemplateID("applicationTemplateId")

read, err := client.GetApplicationTemplate(ctx, id, applicationtemplate.DefaultGetApplicationTemplateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationTemplateClient.GetsCount`

```go
ctx := context.TODO()


read, err := client.GetsCount(ctx, applicationtemplate.DefaultGetsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationTemplateClient.Instantiate`

```go
ctx := context.TODO()
id := applicationtemplate.NewApplicationTemplateID("applicationTemplateId")

payload := applicationtemplate.InstantiateRequest{
	// ...
}


read, err := client.Instantiate(ctx, id, payload, applicationtemplate.DefaultInstantiateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationTemplateClient.ListApplicationTemplates`

```go
ctx := context.TODO()


// alternatively `client.ListApplicationTemplates(ctx, applicationtemplate.DefaultListApplicationTemplatesOperationOptions())` can be used to do batched pagination
items, err := client.ListApplicationTemplatesComplete(ctx, applicationtemplate.DefaultListApplicationTemplatesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
