
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/user` Documentation

The `user` SDK allows for interaction with Microsoft Graph `users` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/user"
```


### Client Initialization

```go
client := user.NewUserClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `UserClient.AssignLicense`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.AssignLicenseRequest{
	// ...
}


read, err := client.AssignLicense(ctx, id, payload, user.DefaultAssignLicenseOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.ChangePassword`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.ChangePasswordRequest{
	// ...
}


read, err := client.ChangePassword(ctx, id, payload, user.DefaultChangePasswordOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, user.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, user.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, user.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, user.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.CreateConvertExternalToInternalMemberUser`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.CreateConvertExternalToInternalMemberUserRequest{
	// ...
}


read, err := client.CreateConvertExternalToInternalMemberUser(ctx, id, payload, user.DefaultCreateConvertExternalToInternalMemberUserOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateExportPersonalData`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.CreateExportPersonalDataRequest{
	// ...
}


read, err := client.CreateExportPersonalData(ctx, id, payload, user.DefaultCreateExportPersonalDataOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateGetsUserOwnedObject`

```go
ctx := context.TODO()

payload := user.CreateGetsUserOwnedObjectRequest{
	// ...
}


read, err := client.CreateGetsUserOwnedObject(ctx, payload, user.DefaultCreateGetsUserOwnedObjectOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateInvalidateAllRefreshToken`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

read, err := client.CreateInvalidateAllRefreshToken(ctx, id, user.DefaultCreateInvalidateAllRefreshTokenOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUnblockManagedApp`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

read, err := client.CreateUnblockManagedApp(ctx, id, user.DefaultCreateUnblockManagedAppOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateUser`

```go
ctx := context.TODO()

payload := user.User{
	// ...
}


read, err := client.CreateUser(ctx, payload, user.DefaultCreateUserOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateValidatesPassword`

```go
ctx := context.TODO()

payload := user.CreateValidatesPasswordRequest{
	// ...
}


read, err := client.CreateValidatesPassword(ctx, payload, user.DefaultCreateValidatesPasswordOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.CreateValidatesProperty`

```go
ctx := context.TODO()

payload := user.CreateValidatesPropertyRequest{
	// ...
}


read, err := client.CreateValidatesProperty(ctx, payload, user.DefaultCreateValidatesPropertyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.DeletePasswordSingleSignOnCredentials`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.DeletePasswordSingleSignOnCredentialsRequest{
	// ...
}


read, err := client.DeletePasswordSingleSignOnCredentials(ctx, id, payload, user.DefaultDeletePasswordSingleSignOnCredentialsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.DeleteUser`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

read, err := client.DeleteUser(ctx, id, user.DefaultDeleteUserOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.FindMeetingTimes`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.FindMeetingTimesRequest{
	// ...
}


read, err := client.FindMeetingTimes(ctx, id, payload, user.DefaultFindMeetingTimesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetMailTips`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.GetMailTipsRequest{
	// ...
}


// alternatively `client.GetMailTips(ctx, id, payload, user.DefaultGetMailTipsOperationOptions())` can be used to do batched pagination
items, err := client.GetMailTipsComplete(ctx, id, payload, user.DefaultGetMailTipsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.GetMemberGroups`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, user.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, user.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.GetMemberObjects`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, user.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, user.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.GetPasswordSingleSignOnCredentials`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

// alternatively `client.GetPasswordSingleSignOnCredentials(ctx, id, user.DefaultGetPasswordSingleSignOnCredentialsOperationOptions())` can be used to do batched pagination
items, err := client.GetPasswordSingleSignOnCredentialsComplete(ctx, id, user.DefaultGetPasswordSingleSignOnCredentialsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.GetUser`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

read, err := client.GetUser(ctx, id, user.DefaultGetUserOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.GetsCount`

```go
ctx := context.TODO()


read, err := client.GetsCount(ctx, user.DefaultGetsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.ListGetsByIds`

```go
ctx := context.TODO()

payload := user.ListGetsByIdsRequest{
	// ...
}


// alternatively `client.ListGetsByIds(ctx, payload, user.DefaultListGetsByIdsOperationOptions())` can be used to do batched pagination
items, err := client.ListGetsByIdsComplete(ctx, payload, user.DefaultListGetsByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.ListUsers`

```go
ctx := context.TODO()


// alternatively `client.ListUsers(ctx, user.DefaultListUsersOperationOptions())` can be used to do batched pagination
items, err := client.ListUsersComplete(ctx, user.DefaultListUsersOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.RemoveAllDevicesFromManagement`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

read, err := client.RemoveAllDevicesFromManagement(ctx, id, user.DefaultRemoveAllDevicesFromManagementOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.ReprocessLicenseAssignment`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

read, err := client.ReprocessLicenseAssignment(ctx, id, user.DefaultReprocessLicenseAssignmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.Restore`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

read, err := client.Restore(ctx, id, user.DefaultRestoreOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.RetryServiceProvisioning`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

read, err := client.RetryServiceProvisioning(ctx, id, user.DefaultRetryServiceProvisioningOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.RevokeSignInSessions`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

read, err := client.RevokeSignInSessions(ctx, id, user.DefaultRevokeSignInSessionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.SendMail`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.SendMailRequest{
	// ...
}


read, err := client.SendMail(ctx, id, payload, user.DefaultSendMailOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.TranslateExchangeIds`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.TranslateExchangeIdsRequest{
	// ...
}


// alternatively `client.TranslateExchangeIds(ctx, id, payload, user.DefaultTranslateExchangeIdsOperationOptions())` can be used to do batched pagination
items, err := client.TranslateExchangeIdsComplete(ctx, id, payload, user.DefaultTranslateExchangeIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `UserClient.UpdateUser`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.User{
	// ...
}


read, err := client.UpdateUser(ctx, id, payload, user.DefaultUpdateUserOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.WipeAndBlockManagedApps`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

read, err := client.WipeAndBlockManagedApps(ctx, id, user.DefaultWipeAndBlockManagedAppsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.WipeManagedAppRegistrationByDeviceTag`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.WipeManagedAppRegistrationByDeviceTagRequest{
	// ...
}


read, err := client.WipeManagedAppRegistrationByDeviceTag(ctx, id, payload, user.DefaultWipeManagedAppRegistrationByDeviceTagOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.WipeManagedAppRegistrationsByAzureAdDeviceId`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.WipeManagedAppRegistrationsByAzureAdDeviceIdRequest{
	// ...
}


read, err := client.WipeManagedAppRegistrationsByAzureAdDeviceId(ctx, id, payload, user.DefaultWipeManagedAppRegistrationsByAzureAdDeviceIdOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `UserClient.WipeManagedAppRegistrationsByDeviceTag`

```go
ctx := context.TODO()
id := user.NewUserID("userId")

payload := user.WipeManagedAppRegistrationsByDeviceTagRequest{
	// ...
}


read, err := client.WipeManagedAppRegistrationsByDeviceTag(ctx, id, payload, user.DefaultWipeManagedAppRegistrationsByDeviceTagOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
