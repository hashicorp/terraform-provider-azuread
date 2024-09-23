
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryobjects/stable/directoryobject` Documentation

The `directoryobject` SDK allows for interaction with Microsoft Graph `directoryobjects` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryobjects/stable/directoryobject"
```


### Client Initialization

```go
client := directoryobject.NewDirectoryObjectClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryObjectClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectId")

payload := directoryobject.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, directoryobject.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, directoryobject.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectId")

payload := directoryobject.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, directoryobject.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, directoryobject.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.CreateDirectoryObject`

```go
ctx := context.TODO()

payload := directoryobject.DirectoryObject{
	// ...
}


read, err := client.CreateDirectoryObject(ctx, payload, directoryobject.DefaultCreateDirectoryObjectOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.CreateValidatesProperty`

```go
ctx := context.TODO()

payload := directoryobject.CreateValidatesPropertyRequest{
	// ...
}


read, err := client.CreateValidatesProperty(ctx, payload, directoryobject.DefaultCreateValidatesPropertyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.DeleteDirectoryObject`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectId")

read, err := client.DeleteDirectoryObject(ctx, id, directoryobject.DefaultDeleteDirectoryObjectOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.GetDirectoryObject`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectId")

read, err := client.GetDirectoryObject(ctx, id, directoryobject.DefaultGetDirectoryObjectOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.GetMemberGroups`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectId")

payload := directoryobject.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, directoryobject.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, directoryobject.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.GetMemberObjects`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectId")

payload := directoryobject.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, directoryobject.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, directoryobject.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.GetsCount`

```go
ctx := context.TODO()


read, err := client.GetsCount(ctx, directoryobject.DefaultGetsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.ListDirectoryObjects`

```go
ctx := context.TODO()


// alternatively `client.ListDirectoryObjects(ctx, directoryobject.DefaultListDirectoryObjectsOperationOptions())` can be used to do batched pagination
items, err := client.ListDirectoryObjectsComplete(ctx, directoryobject.DefaultListDirectoryObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.ListGetsAvailableExtensionProperties`

```go
ctx := context.TODO()

payload := directoryobject.ListGetsAvailableExtensionPropertiesRequest{
	// ...
}


// alternatively `client.ListGetsAvailableExtensionProperties(ctx, payload, directoryobject.DefaultListGetsAvailableExtensionPropertiesOperationOptions())` can be used to do batched pagination
items, err := client.ListGetsAvailableExtensionPropertiesComplete(ctx, payload, directoryobject.DefaultListGetsAvailableExtensionPropertiesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.ListGetsByIds`

```go
ctx := context.TODO()

payload := directoryobject.ListGetsByIdsRequest{
	// ...
}


// alternatively `client.ListGetsByIds(ctx, payload, directoryobject.DefaultListGetsByIdsOperationOptions())` can be used to do batched pagination
items, err := client.ListGetsByIdsComplete(ctx, payload, directoryobject.DefaultListGetsByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryObjectClient.Restore`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectId")

read, err := client.Restore(ctx, id, directoryobject.DefaultRestoreOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryObjectClient.UpdateDirectoryObject`

```go
ctx := context.TODO()
id := directoryobject.NewDirectoryObjectID("directoryObjectId")

payload := directoryobject.DirectoryObject{
	// ...
}


read, err := client.UpdateDirectoryObject(ctx, id, payload, directoryobject.DefaultUpdateDirectoryObjectOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
