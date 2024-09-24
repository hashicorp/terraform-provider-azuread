
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group` Documentation

The `group` SDK allows for interaction with Microsoft Graph `groups` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/groups/beta/group"
```


### Client Initialization

```go
client := group.NewGroupClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `GroupClient.AddFavorite`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

read, err := client.AddFavorite(ctx, id, group.DefaultAddFavoriteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.AssignLicense`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

payload := group.AssignLicenseRequest{
	// ...
}


read, err := client.AssignLicense(ctx, id, payload, group.DefaultAssignLicenseOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CheckGrantedPermissionsForApps`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

// alternatively `client.CheckGrantedPermissionsForApps(ctx, id, group.DefaultCheckGrantedPermissionsForAppsOperationOptions())` can be used to do batched pagination
items, err := client.CheckGrantedPermissionsForAppsComplete(ctx, id, group.DefaultCheckGrantedPermissionsForAppsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

payload := group.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, group.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, group.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

payload := group.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, group.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, group.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.CreateEvaluatesDynamicMembership`

```go
ctx := context.TODO()

payload := group.CreateEvaluatesDynamicMembershipRequest{
	// ...
}


read, err := client.CreateEvaluatesDynamicMembership(ctx, payload, group.DefaultCreateEvaluatesDynamicMembershipOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CreateGetsUserOwnedObject`

```go
ctx := context.TODO()

payload := group.CreateGetsUserOwnedObjectRequest{
	// ...
}


read, err := client.CreateGetsUserOwnedObject(ctx, payload, group.DefaultCreateGetsUserOwnedObjectOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CreateGroup`

```go
ctx := context.TODO()

payload := group.Group{
	// ...
}


read, err := client.CreateGroup(ctx, payload, group.DefaultCreateGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CreateSubscribeByMail`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

read, err := client.CreateSubscribeByMail(ctx, id, group.DefaultCreateSubscribeByMailOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CreateUnsubscribeByMail`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

read, err := client.CreateUnsubscribeByMail(ctx, id, group.DefaultCreateUnsubscribeByMailOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.CreateValidatesProperty`

```go
ctx := context.TODO()

payload := group.CreateValidatesPropertyRequest{
	// ...
}


read, err := client.CreateValidatesProperty(ctx, payload, group.DefaultCreateValidatesPropertyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.DeleteGroup`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

read, err := client.DeleteGroup(ctx, id, group.DefaultDeleteGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.DeletePasswordSingleSignOnCredentials`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

payload := group.DeletePasswordSingleSignOnCredentialsRequest{
	// ...
}


read, err := client.DeletePasswordSingleSignOnCredentials(ctx, id, payload, group.DefaultDeletePasswordSingleSignOnCredentialsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.EvaluateDynamicMembership`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

payload := group.EvaluateDynamicMembershipRequest{
	// ...
}


read, err := client.EvaluateDynamicMembership(ctx, id, payload, group.DefaultEvaluateDynamicMembershipOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.GetGroup`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

read, err := client.GetGroup(ctx, id, group.DefaultGetGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.GetMemberGroups`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

payload := group.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, group.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, group.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.GetMemberObjects`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

payload := group.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, group.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, group.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.GetPasswordSingleSignOnCredentials`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

// alternatively `client.GetPasswordSingleSignOnCredentials(ctx, id, group.DefaultGetPasswordSingleSignOnCredentialsOperationOptions())` can be used to do batched pagination
items, err := client.GetPasswordSingleSignOnCredentialsComplete(ctx, id, group.DefaultGetPasswordSingleSignOnCredentialsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.GetsCount`

```go
ctx := context.TODO()


read, err := client.GetsCount(ctx, group.DefaultGetsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.ListGetsByIds`

```go
ctx := context.TODO()

payload := group.ListGetsByIdsRequest{
	// ...
}


// alternatively `client.ListGetsByIds(ctx, payload, group.DefaultListGetsByIdsOperationOptions())` can be used to do batched pagination
items, err := client.ListGetsByIdsComplete(ctx, payload, group.DefaultListGetsByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.ListGroups`

```go
ctx := context.TODO()


// alternatively `client.ListGroups(ctx, group.DefaultListGroupsOperationOptions())` can be used to do batched pagination
items, err := client.ListGroupsComplete(ctx, group.DefaultListGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `GroupClient.RemoveFavorite`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

read, err := client.RemoveFavorite(ctx, id, group.DefaultRemoveFavoriteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.Renew`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

read, err := client.Renew(ctx, id, group.DefaultRenewOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.ResetUnseenCount`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

read, err := client.ResetUnseenCount(ctx, id, group.DefaultResetUnseenCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.Restore`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

payload := group.RestoreRequest{
	// ...
}


read, err := client.Restore(ctx, id, payload, group.DefaultRestoreOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.RetryServiceProvisioning`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

read, err := client.RetryServiceProvisioning(ctx, id, group.DefaultRetryServiceProvisioningOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.UpdateGroup`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

payload := group.Group{
	// ...
}


read, err := client.UpdateGroup(ctx, id, payload, group.DefaultUpdateGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `GroupClient.ValidateProperties`

```go
ctx := context.TODO()
id := group.NewGroupID("groupId")

payload := group.ValidatePropertiesRequest{
	// ...
}


read, err := client.ValidateProperties(ctx, id, payload, group.DefaultValidatePropertiesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
