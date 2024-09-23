
## `github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/stable/domain` Documentation

The `domain` SDK allows for interaction with Microsoft Graph `domains` (API Version `stable`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/stable/domain"
```


### Client Initialization

```go
client := domain.NewDomainClientWithBaseURI("https://graph.microsoft.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DomainClient.CreateDomain`

```go
ctx := context.TODO()

payload := domain.Domain{
	// ...
}


read, err := client.CreateDomain(ctx, payload, domain.DefaultCreateDomainOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.CreateForceDelete`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainId")

payload := domain.CreateForceDeleteRequest{
	// ...
}


read, err := client.CreateForceDelete(ctx, id, payload, domain.DefaultCreateForceDeleteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.CreatePromote`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainId")

read, err := client.CreatePromote(ctx, id, domain.DefaultCreatePromoteOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.CreateVerify`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainId")

read, err := client.CreateVerify(ctx, id, domain.DefaultCreateVerifyOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.DeleteDomain`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainId")

read, err := client.DeleteDomain(ctx, id, domain.DefaultDeleteDomainOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.GetDomain`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainId")

read, err := client.GetDomain(ctx, id, domain.DefaultGetDomainOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.GetsCount`

```go
ctx := context.TODO()


read, err := client.GetsCount(ctx, domain.DefaultGetsCountOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `DomainClient.ListDomains`

```go
ctx := context.TODO()


// alternatively `client.ListDomains(ctx, domain.DefaultListDomainsOperationOptions())` can be used to do batched pagination
items, err := client.ListDomainsComplete(ctx, domain.DefaultListDomainsOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `DomainClient.UpdateDomain`

```go
ctx := context.TODO()
id := domain.NewDomainID("domainId")

payload := domain.Domain{
	// ...
}


read, err := client.UpdateDomain(ctx, id, payload, domain.DefaultUpdateDomainOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
