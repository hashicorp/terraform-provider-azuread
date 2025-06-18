
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipal` Documentation

The `serviceprincipal` SDK allows for interaction with Microsoft Graph `serviceprincipals` (API Version `beta`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipal"
```


### Client Initialization

```go
client := serviceprincipal.NewServicePrincipalClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServicePrincipalClient.AddTokenSigningCertificate`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

payload := serviceprincipal.AddTokenSigningCertificateRequest{
	// ...
}


read, err := client.AddTokenSigningCertificate(ctx, id, payload, serviceprincipal.DefaultAddTokenSigningCertificateOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.CheckMemberGroups`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

payload := serviceprincipal.CheckMemberGroupsRequest{
	// ...
}


// alternatively `client.CheckMemberGroups(ctx, id, payload, serviceprincipal.DefaultCheckMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberGroupsComplete(ctx, id, payload, serviceprincipal.DefaultCheckMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalClient.CheckMemberObjects`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

payload := serviceprincipal.CheckMemberObjectsRequest{
	// ...
}


// alternatively `client.CheckMemberObjects(ctx, id, payload, serviceprincipal.DefaultCheckMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.CheckMemberObjectsComplete(ctx, id, payload, serviceprincipal.DefaultCheckMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalClient.CreateGetsUserOwnedObject`

```go
ctx := context.TODO()

payload := serviceprincipal.CreateGetsUserOwnedObjectRequest{
	// ...
}


read, err := client.CreateGetsUserOwnedObject(ctx, payload, serviceprincipal.DefaultCreateGetsUserOwnedObjectOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.CreatePasswordSingleSignOnCredentials`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

payload := serviceprincipal.CreatePasswordSingleSignOnCredentialsRequest{
	// ...
}


read, err := client.CreatePasswordSingleSignOnCredentials(ctx, id, payload, serviceprincipal.DefaultCreatePasswordSingleSignOnCredentialsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.CreateServicePrincipal`

```go
ctx := context.TODO()

payload := serviceprincipal.ServicePrincipal{
	// ...
}


read, err := client.CreateServicePrincipal(ctx, payload, serviceprincipal.DefaultCreateServicePrincipalOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.CreateValidatesProperty`

```go
ctx := context.TODO()

payload := serviceprincipal.CreateValidatesPropertyRequest{
	// ...
}


read, err := client.CreateValidatesProperty(ctx, payload, serviceprincipal.DefaultCreateValidatesPropertyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.DeletePasswordSingleSignOnCredentials`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

payload := serviceprincipal.DeletePasswordSingleSignOnCredentialsRequest{
	// ...
}


read, err := client.DeletePasswordSingleSignOnCredentials(ctx, id, payload, serviceprincipal.DefaultDeletePasswordSingleSignOnCredentialsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.DeleteServicePrincipal`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

read, err := client.DeleteServicePrincipal(ctx, id, serviceprincipal.DefaultDeleteServicePrincipalOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetMemberGroups`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

payload := serviceprincipal.GetMemberGroupsRequest{
	// ...
}


// alternatively `client.GetMemberGroups(ctx, id, payload, serviceprincipal.DefaultGetMemberGroupsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberGroupsComplete(ctx, id, payload, serviceprincipal.DefaultGetMemberGroupsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalClient.GetMemberObjects`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

payload := serviceprincipal.GetMemberObjectsRequest{
	// ...
}


// alternatively `client.GetMemberObjects(ctx, id, payload, serviceprincipal.DefaultGetMemberObjectsOperationOptions())` can be used to do batched pagination
items, err := client.GetMemberObjectsComplete(ctx, id, payload, serviceprincipal.DefaultGetMemberObjectsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalClient.GetPasswordSingleSignOnCredentials`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

payload := serviceprincipal.GetPasswordSingleSignOnCredentialsRequest{
	// ...
}


read, err := client.GetPasswordSingleSignOnCredentials(ctx, id, payload, serviceprincipal.DefaultGetPasswordSingleSignOnCredentialsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetServicePrincipal`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

read, err := client.GetServicePrincipal(ctx, id, serviceprincipal.DefaultGetServicePrincipalOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.GetsCount`

```go
ctx := context.TODO()


read, err := client.GetsCount(ctx, serviceprincipal.DefaultGetsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.ListGetsByIds`

```go
ctx := context.TODO()

payload := serviceprincipal.ListGetsByIdsRequest{
	// ...
}


// alternatively `client.ListGetsByIds(ctx, payload, serviceprincipal.DefaultListGetsByIdsOperationOptions())` can be used to do batched pagination
items, err := client.ListGetsByIdsComplete(ctx, payload, serviceprincipal.DefaultListGetsByIdsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalClient.ListServicePrincipals`

```go
ctx := context.TODO()


// alternatively `client.ListServicePrincipals(ctx, serviceprincipal.DefaultListServicePrincipalsOperationOptions())` can be used to do batched pagination
items, err := client.ListServicePrincipalsComplete(ctx, serviceprincipal.DefaultListServicePrincipalsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `ServicePrincipalClient.Restore`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

read, err := client.Restore(ctx, id, serviceprincipal.DefaultRestoreOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.UpdatePasswordSingleSignOnCredentials`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

payload := serviceprincipal.UpdatePasswordSingleSignOnCredentialsRequest{
	// ...
}


read, err := client.UpdatePasswordSingleSignOnCredentials(ctx, id, payload, serviceprincipal.DefaultUpdatePasswordSingleSignOnCredentialsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.UpdateServicePrincipal`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

payload := serviceprincipal.ServicePrincipal{
	// ...
}


read, err := client.UpdateServicePrincipal(ctx, id, payload, serviceprincipal.DefaultUpdateServicePrincipalOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
