
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/synchronizationsecret` Documentation

The `synchronizationsecret` SDK allows for interaction with Microsoft Graph `serviceprincipals` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/synchronizationsecret"
```


### Client Initialization

```go
client := synchronizationsecret.NewSynchronizationSecretClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SynchronizationSecretClient.GetSynchronizationSecretsCount`

```go
ctx := context.TODO()
id := synchronizationsecret.NewServicePrincipalID("servicePrincipalId")

read, err := client.GetSynchronizationSecretsCount(ctx, id, synchronizationsecret.DefaultGetSynchronizationSecretsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationSecretClient.ListSynchronizationSecrets`

```go
ctx := context.TODO()
id := synchronizationsecret.NewServicePrincipalID("servicePrincipalId")

// alternatively `client.ListSynchronizationSecrets(ctx, id, synchronizationsecret.DefaultListSynchronizationSecretsOperationOptions())` can be used to do batched pagination
items, err := client.ListSynchronizationSecretsComplete(ctx, id, synchronizationsecret.DefaultListSynchronizationSecretsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SynchronizationSecretClient.SetSynchronizationSecret`

```go
ctx := context.TODO()
id := synchronizationsecret.NewServicePrincipalID("servicePrincipalId")

payload := synchronizationsecret.SetSynchronizationSecretRequest{
	// ...
}


read, err := client.SetSynchronizationSecret(ctx, id, payload, synchronizationsecret.DefaultSetSynchronizationSecretOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
