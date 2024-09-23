
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroletemplates/stable/directoryroletemplate` Documentation

The `directoryroletemplate` SDK allows for interaction with Microsoft Graph `directoryroletemplates` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroletemplates/stable/directoryroletemplate"
```


### Client Initialization

```go
client := directoryroletemplate.NewDirectoryRoleTemplateClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DirectoryRoleTemplateClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateId")

payload := directoryroletemplate.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, directoryroletemplate.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, directoryroletemplate.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleTemplateClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateId")

payload := directoryroletemplate.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, directoryroletemplate.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, directoryroletemplate.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleTemplateClient.CreateDirectoryRoleTemplate`

```go
ctx := context.TODO()

payload := directoryroletemplate.DirectoryRoleTemplate{
	// ...
}


read, err := client.CreateDirectoryRoleTemplate(ctx, payload, directoryroletemplate.DefaultCreateDirectoryRoleTemplateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleTemplateClient.CreateValidatesProperty`

```go
ctx := context.TODO()

payload := directoryroletemplate.CreateValidatesPropertyRequest{
	// ...
}


read, err := client.CreateValidatesProperty(ctx, payload, directoryroletemplate.DefaultCreateValidatesPropertyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleTemplateClient.DeleteDirectoryRoleTemplate`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateId")

read, err := client.DeleteDirectoryRoleTemplate(ctx, id, directoryroletemplate.DefaultDeleteDirectoryRoleTemplateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleTemplateClient.GetDirectoryRoleTemplate`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateId")

read, err := client.GetDirectoryRoleTemplate(ctx, id, directoryroletemplate.DefaultGetDirectoryRoleTemplateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleTemplateClient.GetMemberGroups`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateId")

payload := directoryroletemplate.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, directoryroletemplate.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, directoryroletemplate.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleTemplateClient.GetMemberObjects`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateId")

payload := directoryroletemplate.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, directoryroletemplate.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, directoryroletemplate.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleTemplateClient.GetsCount`

```go
ctx := context.TODO()


read, err := client.GetsCount(ctx, directoryroletemplate.DefaultGetsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleTemplateClient.ListDirectoryRoleTemplates`

```go
ctx := context.TODO()


// alternatively `client.ListDirectoryRoleTemplates(ctx, directoryroletemplate.DefaultListDirectoryRoleTemplatesOperationOptions())` can be used to do batched pagination
items, err := client.ListDirectoryRoleTemplatesComplete(ctx, directoryroletemplate.DefaultListDirectoryRoleTemplatesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleTemplateClient.ListGetsAvailableExtensionProperties`

```go
ctx := context.TODO()

payload := directoryroletemplate.ListGetsAvailableExtensionPropertiesRequest{
	// ...
}


// alternatively `client.ListGetsAvailableExtensionProperties(ctx, payload, directoryroletemplate.DefaultListGetsAvailableExtensionPropertiesOperationOptions())` can be used to do batched pagination
items, err := client.ListGetsAvailableExtensionPropertiesComplete(ctx, payload, directoryroletemplate.DefaultListGetsAvailableExtensionPropertiesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleTemplateClient.ListGetsByIds`

```go
ctx := context.TODO()

payload := directoryroletemplate.ListGetsByIdsRequest{
	// ...
}


// alternatively `client.ListGetsByIds(ctx, payload, directoryroletemplate.DefaultListGetsByIdsOperationOptions())` can be used to do batched pagination
items, err := client.ListGetsByIdsComplete(ctx, payload, directoryroletemplate.DefaultListGetsByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DirectoryRoleTemplateClient.Restore`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateId")

read, err := client.Restore(ctx, id, directoryroletemplate.DefaultRestoreOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DirectoryRoleTemplateClient.UpdateDirectoryRoleTemplate`

```go
ctx := context.TODO()
id := directoryroletemplate.NewDirectoryRoleTemplateID("directoryRoleTemplateId")

payload := directoryroletemplate.DirectoryRoleTemplate{
	// ...
}


read, err := client.UpdateDirectoryRoleTemplate(ctx, id, payload, directoryroletemplate.DefaultUpdateDirectoryRoleTemplateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
