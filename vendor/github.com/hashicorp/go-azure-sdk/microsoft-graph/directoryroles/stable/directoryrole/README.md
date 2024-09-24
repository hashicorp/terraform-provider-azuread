
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroles/stable/directoryrole` Documentation

The `directoryrole` SDK allows for interaction with Microsoft Graph `directoryroles` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroles/stable/directoryrole"
```


### Client Initialization

```go
client := directoryrole.NewDirectoryRoleClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryRoleClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleId")

payload := directoryrole.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, directoryrole.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, directoryrole.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleId")

payload := directoryrole.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, directoryrole.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, directoryrole.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleClient.CreateDirectoryRole`

```go
ctx := context.TODO()

payload := directoryrole.DirectoryRole{
	// ...
}


read, err := client.CreateDirectoryRole(ctx, payload, directoryrole.DefaultCreateDirectoryRoleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleClient.CreateValidatesProperty`

```go
ctx := context.TODO()

payload := directoryrole.CreateValidatesPropertyRequest{
	// ...
}


read, err := client.CreateValidatesProperty(ctx, payload, directoryrole.DefaultCreateValidatesPropertyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleClient.DeleteDirectoryRole`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleId")

read, err := client.DeleteDirectoryRole(ctx, id, directoryrole.DefaultDeleteDirectoryRoleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleClient.GetDirectoryRole`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleId")

read, err := client.GetDirectoryRole(ctx, id, directoryrole.DefaultGetDirectoryRoleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleClient.GetMemberGroups`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleId")

payload := directoryrole.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, directoryrole.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, directoryrole.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleClient.GetMemberObjects`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleId")

payload := directoryrole.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, directoryrole.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, directoryrole.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleClient.GetsCount`

```go
ctx := context.TODO()


read, err := client.GetsCount(ctx, directoryrole.DefaultGetsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleClient.ListDirectoryRoles`

```go
ctx := context.TODO()


// alternatively `client.ListDirectoryRoles(ctx, directoryrole.DefaultListDirectoryRolesOperationOptions())` can be used to do batched pagination
items, err := client.ListDirectoryRolesComplete(ctx, directoryrole.DefaultListDirectoryRolesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleClient.ListGetsAvailableExtensionProperties`

```go
ctx := context.TODO()

payload := directoryrole.ListGetsAvailableExtensionPropertiesRequest{
	// ...
}


// alternatively `client.ListGetsAvailableExtensionProperties(ctx, payload, directoryrole.DefaultListGetsAvailableExtensionPropertiesOperationOptions())` can be used to do batched pagination
items, err := client.ListGetsAvailableExtensionPropertiesComplete(ctx, payload, directoryrole.DefaultListGetsAvailableExtensionPropertiesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleClient.ListGetsByIds`

```go
ctx := context.TODO()

payload := directoryrole.ListGetsByIdsRequest{
	// ...
}


// alternatively `client.ListGetsByIds(ctx, payload, directoryrole.DefaultListGetsByIdsOperationOptions())` can be used to do batched pagination
items, err := client.ListGetsByIdsComplete(ctx, payload, directoryrole.DefaultListGetsByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleClient.Restore`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleId")

read, err := client.Restore(ctx, id, directoryrole.DefaultRestoreOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleClient.UpdateDirectoryRole`

```go
ctx := context.TODO()
id := directoryrole.NewDirectoryRoleID("directoryRoleId")

payload := directoryrole.DirectoryRole{
	// ...
}


read, err := client.UpdateDirectoryRole(ctx, id, payload, directoryrole.DefaultUpdateDirectoryRoleOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
