
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/synchronizationjob` Documentation

The `synchronizationjob` SDK allows for interaction with Microsoft Graph `serviceprincipals` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/synchronizationjob"
```


### Client Initialization

```go
client := synchronizationjob.NewSynchronizationJobClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SynchronizationJobClient.CreateSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalID("servicePrincipalId")

payload := synchronizationjob.SynchronizationJob{
	// ...
}


read, err := client.CreateSynchronizationJob(ctx, id, payload, synchronizationjob.DefaultCreateSynchronizationJobOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.DeleteSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalId", "synchronizationJobId")

read, err := client.DeleteSynchronizationJob(ctx, id, synchronizationjob.DefaultDeleteSynchronizationJobOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.GetSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalId", "synchronizationJobId")

read, err := client.GetSynchronizationJob(ctx, id, synchronizationjob.DefaultGetSynchronizationJobOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.GetSynchronizationJobsCount`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalID("servicePrincipalId")

read, err := client.GetSynchronizationJobsCount(ctx, id, synchronizationjob.DefaultGetSynchronizationJobsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.ListSynchronizationJobs`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalID("servicePrincipalId")

// alternatively `client.ListSynchronizationJobs(ctx, id, synchronizationjob.DefaultListSynchronizationJobsOperationOptions())` can be used to do batched pagination
items, err := client.ListSynchronizationJobsComplete(ctx, id, synchronizationjob.DefaultListSynchronizationJobsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `SynchronizationJobClient.PauseSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalId", "synchronizationJobId")

read, err := client.PauseSynchronizationJob(ctx, id, synchronizationjob.DefaultPauseSynchronizationJobOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.ProvisionSynchronizationJobOnDemand`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalId", "synchronizationJobId")

payload := synchronizationjob.ProvisionSynchronizationJobOnDemandRequest{
	// ...
}


read, err := client.ProvisionSynchronizationJobOnDemand(ctx, id, payload, synchronizationjob.DefaultProvisionSynchronizationJobOnDemandOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.RestartSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalId", "synchronizationJobId")

payload := synchronizationjob.RestartSynchronizationJobRequest{
	// ...
}


read, err := client.RestartSynchronizationJob(ctx, id, payload, synchronizationjob.DefaultRestartSynchronizationJobOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.StartSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalId", "synchronizationJobId")

read, err := client.StartSynchronizationJob(ctx, id, synchronizationjob.DefaultStartSynchronizationJobOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.UpdateSynchronizationJob`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalId", "synchronizationJobId")

payload := synchronizationjob.SynchronizationJob{
	// ...
}


read, err := client.UpdateSynchronizationJob(ctx, id, payload, synchronizationjob.DefaultUpdateSynchronizationJobOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.ValidateSynchronizationJobCredentials`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalIdSynchronizationJobID("servicePrincipalId", "synchronizationJobId")

payload := synchronizationjob.ValidateSynchronizationJobCredentialsRequest{
	// ...
}


read, err := client.ValidateSynchronizationJobCredentials(ctx, id, payload, synchronizationjob.DefaultValidateSynchronizationJobCredentialsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `SynchronizationJobClient.ValidateSynchronizationJobsCredentials`

```go
ctx := context.TODO()
id := synchronizationjob.NewServicePrincipalID("servicePrincipalId")

payload := synchronizationjob.ValidateSynchronizationJobsCredentialsRequest{
	// ...
}


read, err := client.ValidateSynchronizationJobsCredentials(ctx, id, payload, synchronizationjob.DefaultValidateSynchronizationJobsCredentialsOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
