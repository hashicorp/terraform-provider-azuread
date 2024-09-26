
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/administrativeunits/beta/administrativeunit` Documentation

The `administrativeunit` SDK allows for interaction with Microsoft Graph `administrativeunits` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/administrativeunits/beta/administrativeunit"
```


### Client Initialization

```go
client := administrativeunit.NewAdministrativeUnitClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `AdministrativeUnitClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitId")

payload := administrativeunit.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, administrativeunit.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, administrativeunit.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdministrativeUnitClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitId")

payload := administrativeunit.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, administrativeunit.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, administrativeunit.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdministrativeUnitClient.CreateAdministrativeUnit`

```go
ctx := context.TODO()

payload := administrativeunit.AdministrativeUnit{
	// ...
}


read, err := client.CreateAdministrativeUnit(ctx, payload, administrativeunit.DefaultCreateAdministrativeUnitOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.CreateGetsUserOwnedObject`

```go
ctx := context.TODO()

payload := administrativeunit.CreateGetsUserOwnedObjectRequest{
	// ...
}


read, err := client.CreateGetsUserOwnedObject(ctx, payload, administrativeunit.DefaultCreateGetsUserOwnedObjectOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.CreateValidatesProperty`

```go
ctx := context.TODO()

payload := administrativeunit.CreateValidatesPropertyRequest{
	// ...
}


read, err := client.CreateValidatesProperty(ctx, payload, administrativeunit.DefaultCreateValidatesPropertyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.DeleteAdministrativeUnit`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitId")

read, err := client.DeleteAdministrativeUnit(ctx, id, administrativeunit.DefaultDeleteAdministrativeUnitOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.GetAdministrativeUnit`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitId")

read, err := client.GetAdministrativeUnit(ctx, id, administrativeunit.DefaultGetAdministrativeUnitOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.GetMemberGroups`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitId")

payload := administrativeunit.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, administrativeunit.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, administrativeunit.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdministrativeUnitClient.GetMemberObjects`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitId")

payload := administrativeunit.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, administrativeunit.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, administrativeunit.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdministrativeUnitClient.GetsCount`

```go
ctx := context.TODO()


read, err := client.GetsCount(ctx, administrativeunit.DefaultGetsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.ListAdministrativeUnits`

```go
ctx := context.TODO()


// alternatively `client.ListAdministrativeUnits(ctx, administrativeunit.DefaultListAdministrativeUnitsOperationOptions())` can be used to do batched pagination
items, err := client.ListAdministrativeUnitsComplete(ctx, administrativeunit.DefaultListAdministrativeUnitsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdministrativeUnitClient.ListGetsByIds`

```go
ctx := context.TODO()

payload := administrativeunit.ListGetsByIdsRequest{
	// ...
}


// alternatively `client.ListGetsByIds(ctx, payload, administrativeunit.DefaultListGetsByIdsOperationOptions())` can be used to do batched pagination
items, err := client.ListGetsByIdsComplete(ctx, payload, administrativeunit.DefaultListGetsByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `AdministrativeUnitClient.Restore`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitId")

payload := administrativeunit.RestoreRequest{
	// ...
}


read, err := client.Restore(ctx, id, payload, administrativeunit.DefaultRestoreOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `AdministrativeUnitClient.UpdateAdministrativeUnit`

```go
ctx := context.TODO()
id := administrativeunit.NewAdministrativeUnitID("administrativeUnitId")

payload := administrativeunit.AdministrativeUnit{
	// ...
}


read, err := client.UpdateAdministrativeUnit(ctx, id, payload, administrativeunit.DefaultUpdateAdministrativeUnitOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
