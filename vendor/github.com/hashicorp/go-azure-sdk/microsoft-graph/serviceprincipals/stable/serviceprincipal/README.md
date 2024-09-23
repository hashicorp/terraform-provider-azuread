
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal` Documentation

The `serviceprincipal` SDK allows for interaction with Microsoft Graph `serviceprincipals` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
```


### Client Initialization

```go
client := serviceprincipal.NewServicePrincipalClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServicePrincipalClient.AddKey`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

payload := serviceprincipal.AddKeyRequest{
	// ...
}


read, err := client.AddKey(ctx, id, payload, serviceprincipal.DefaultAddKeyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.AddPassword`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

payload := serviceprincipal.AddPasswordRequest{
	// ...
}


read, err := client.AddPassword(ctx, id, payload, serviceprincipal.DefaultAddPasswordOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
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


### Example Usage: `ServicePrincipalClient.ListGetsAvailableExtensionProperties`

```go
ctx := context.TODO()

payload := serviceprincipal.ListGetsAvailableExtensionPropertiesRequest{
	// ...
}


// alternatively `client.ListGetsAvailableExtensionProperties(ctx, payload, serviceprincipal.DefaultListGetsAvailableExtensionPropertiesOperationOptions())` can be used to do batched pagination
items, err := client.ListGetsAvailableExtensionPropertiesComplete(ctx, payload, serviceprincipal.DefaultListGetsAvailableExtensionPropertiesOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
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


### Example Usage: `ServicePrincipalClient.RemoveKey`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

payload := serviceprincipal.RemoveKeyRequest{
	// ...
}


read, err := client.RemoveKey(ctx, id, payload, serviceprincipal.DefaultRemoveKeyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServicePrincipalClient.RemovePassword`

```go
ctx := context.TODO()
id := serviceprincipal.NewServicePrincipalID("servicePrincipalId")

payload := serviceprincipal.RemovePasswordRequest{
	// ...
}


read, err := client.RemovePassword(ctx, id, payload, serviceprincipal.DefaultRemovePasswordOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
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
