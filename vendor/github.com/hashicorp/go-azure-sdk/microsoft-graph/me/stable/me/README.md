
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/me` Documentation

The `me` SDK allows for interaction with Microsoft Graph `me` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/me"
```


### Client Initialization

```go
client := me.NewMeClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MeClient.AssignLicense`

```go
ctx := context.TODO()

payload := me.AssignLicenseRequest{
	// ...
}


read, err := client.AssignLicense(ctx, payload, me.DefaultAssignLicenseOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.ChangePassword`

```go
ctx := context.TODO()

payload := me.ChangePasswordRequest{
	// ...
}


read, err := client.ChangePassword(ctx, payload, me.DefaultChangePasswordOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.CheckMemberGroups`

```go
ctx := context.TODO()

payload := me.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, payload, me.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, payload, me.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeClient.CheckMemberObjects`

```go
ctx := context.TODO()

payload := me.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, payload, me.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, payload, me.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeClient.CreateExportPersonalData`

```go
ctx := context.TODO()

payload := me.CreateExportPersonalDataRequest{
	// ...
}


read, err := client.CreateExportPersonalData(ctx, payload, me.DefaultCreateExportPersonalDataOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.FindMeetingTimes`

```go
ctx := context.TODO()

payload := me.FindMeetingTimesRequest{
	// ...
}


read, err := client.FindMeetingTimes(ctx, payload, me.DefaultFindMeetingTimesOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.GetMailTips`

```go
ctx := context.TODO()

payload := me.GetMailTipsRequest{
	// ...
}


// alternatively `client.GetMailTips(ctx, payload, me.DefaultGetMailTipsOperationOptions())` can be used to do batched pagination
items, err := client.GetMailTipsComplete(ctx, payload, me.DefaultGetMailTipsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeClient.GetMe`

```go
ctx := context.TODO()


read, err := client.GetMe(ctx, me.DefaultGetMeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.GetMemberGroups`

```go
ctx := context.TODO()

payload := me.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, payload, me.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, payload, me.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeClient.GetMemberObjects`

```go
ctx := context.TODO()

payload := me.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, payload, me.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, payload, me.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeClient.RemoveAllDevicesFromManagement`

```go
ctx := context.TODO()


read, err := client.RemoveAllDevicesFromManagement(ctx, me.DefaultRemoveAllDevicesFromManagementOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.ReprocessLicenseAssignment`

```go
ctx := context.TODO()


read, err := client.ReprocessLicenseAssignment(ctx, me.DefaultReprocessLicenseAssignmentOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.Restore`

```go
ctx := context.TODO()


read, err := client.Restore(ctx, me.DefaultRestoreOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.RetryServiceProvisioning`

```go
ctx := context.TODO()


read, err := client.RetryServiceProvisioning(ctx, me.DefaultRetryServiceProvisioningOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.RevokeSignInSessions`

```go
ctx := context.TODO()


read, err := client.RevokeSignInSessions(ctx, me.DefaultRevokeSignInSessionsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.SendMail`

```go
ctx := context.TODO()

payload := me.SendMailRequest{
	// ...
}


read, err := client.SendMail(ctx, payload, me.DefaultSendMailOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.TranslateExchangeIds`

```go
ctx := context.TODO()

payload := me.TranslateExchangeIdsRequest{
	// ...
}


// alternatively `client.TranslateExchangeIds(ctx, payload, me.DefaultTranslateExchangeIdsOperationOptions())` can be used to do batched pagination
items, err := client.TranslateExchangeIdsComplete(ctx, payload, me.DefaultTranslateExchangeIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `MeClient.UpdateMe`

```go
ctx := context.TODO()

payload := me.User{
	// ...
}


read, err := client.UpdateMe(ctx, payload, me.DefaultUpdateMeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MeClient.WipeManagedAppRegistrationsByDeviceTag`

```go
ctx := context.TODO()

payload := me.WipeManagedAppRegistrationsByDeviceTagRequest{
	// ...
}


read, err := client.WipeManagedAppRegistrationsByDeviceTag(ctx, payload, me.DefaultWipeManagedAppRegistrationsByDeviceTagOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
