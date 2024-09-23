
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/application` Documentation

The `application` SDK allows for interaction with Microsoft Graph `applications` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/applications/beta/application"
```


### Client Initialization

```go
client := application.NewApplicationClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ApplicationClient.AddKey`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationId")

payload := application.AddKeyRequest{
	// ...
}


read, err := client.AddKey(ctx, id, payload, application.DefaultAddKeyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.AddPassword`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationId")

payload := application.AddPasswordRequest{
	// ...
}


read, err := client.AddPassword(ctx, id, payload, application.DefaultAddPasswordOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationId")

payload := application.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, application.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, application.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationId")

payload := application.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, application.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, application.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.CreateApplication`

```go
ctx := context.TODO()

payload := application.Application{
	// ...
}


read, err := client.CreateApplication(ctx, payload, application.DefaultCreateApplicationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.CreateGetsUserOwnedObject`

```go
ctx := context.TODO()

payload := application.CreateGetsUserOwnedObjectRequest{
	// ...
}


read, err := client.CreateGetsUserOwnedObject(ctx, payload, application.DefaultCreateGetsUserOwnedObjectOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.CreateValidatesProperty`

```go
ctx := context.TODO()

payload := application.CreateValidatesPropertyRequest{
	// ...
}


read, err := client.CreateValidatesProperty(ctx, payload, application.DefaultCreateValidatesPropertyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.DeleteApplication`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationId")

read, err := client.DeleteApplication(ctx, id, application.DefaultDeleteApplicationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.GetApplication`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationId")

read, err := client.GetApplication(ctx, id, application.DefaultGetApplicationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.GetMemberGroups`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationId")

payload := application.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, application.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, application.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.GetMemberObjects`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationId")

payload := application.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, application.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, application.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.GetsCount`

```go
ctx := context.TODO()


read, err := client.GetsCount(ctx, application.DefaultGetsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.ListApplications`

```go
ctx := context.TODO()


// alternatively `client.ListApplications(ctx, application.DefaultListApplicationsOperationOptions())` can be used to do batched pagination
items, err := client.ListApplicationsComplete(ctx, application.DefaultListApplicationsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.ListGetsByIds`

```go
ctx := context.TODO()

payload := application.ListGetsByIdsRequest{
	// ...
}


// alternatively `client.ListGetsByIds(ctx, payload, application.DefaultListGetsByIdsOperationOptions())` can be used to do batched pagination
items, err := client.ListGetsByIdsComplete(ctx, payload, application.DefaultListGetsByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ApplicationClient.RemoveKey`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationId")

payload := application.RemoveKeyRequest{
	// ...
}


read, err := client.RemoveKey(ctx, id, payload, application.DefaultRemoveKeyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.RemovePassword`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationId")

payload := application.RemovePasswordRequest{
	// ...
}


read, err := client.RemovePassword(ctx, id, payload, application.DefaultRemovePasswordOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.Restore`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationId")

payload := application.RestoreRequest{
	// ...
}


read, err := client.Restore(ctx, id, payload, application.DefaultRestoreOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.SetVerifiedPublisher`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationId")

payload := application.SetVerifiedPublisherRequest{
	// ...
}


read, err := client.SetVerifiedPublisher(ctx, id, payload, application.DefaultSetVerifiedPublisherOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.UnsetVerifiedPublisher`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationId")

read, err := client.UnsetVerifiedPublisher(ctx, id, application.DefaultUnsetVerifiedPublisherOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ApplicationClient.UpdateApplication`

```go
ctx := context.TODO()
id := application.NewApplicationID("applicationId")

payload := application.Application{
	// ...
}


read, err := client.UpdateApplication(ctx, id, payload, application.DefaultUpdateApplicationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
