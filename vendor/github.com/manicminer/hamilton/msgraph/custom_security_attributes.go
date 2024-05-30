package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/manicminer/hamilton/internal/utils"
)

const (
	// customSecurityAttributeDefinitionEntity is a static string used by all methods on the
	// CustomSecurityAttributeDefinitionClient struct
	customSecurityAttributeDefinitionEntity = "/directory/customSecurityAttributeDefinitions"
)

// CustomSecurityAttributeDefinitionClient returns a BaseClient to enable interaction with the
// graph API
type CustomSecurityAttributeDefinitionClient struct {
	BaseClient Client
}

// NewCustomSecurityAttributeDefinitionClient returns a new instance of
// CustomSecurityAttributeDefinitionClient
func NewCustomSecurityAttributeDefinitionClient() *CustomSecurityAttributeDefinitionClient {
	return &CustomSecurityAttributeDefinitionClient{
		BaseClient: NewClient(Version10),
	}
}

// List returns a slice of CustomSecurityAttributeDefinition, the HTTP status code and any errors
func (c *CustomSecurityAttributeDefinitionClient) List(ctx context.Context, query odata.Query) (*[]CustomSecurityAttributeDefinition, int, error) {
	resp, status, _, err := c.BaseClient.Get(
		ctx,
		GetHttpRequestInput{
			OData:            query,
			ValidStatusCodes: []int{http.StatusOK},
			Uri: Uri{
				Entity: customSecurityAttributeDefinitionEntity,
			},
		},
	)
	if err != nil {
		return nil, status, fmt.Errorf("CustomSecurityAttributeDefinition.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		CustomSecurityAttributeDefinitions []CustomSecurityAttributeDefinition `json:"value"`
	}

	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.CustomSecurityAttributeDefinitions, status, nil
}

// Create will create a CustomSecurityAttributeDefinition and return the result, HTTP status code
// as well as any errors
func (c *CustomSecurityAttributeDefinitionClient) Create(ctx context.Context, customSecurityAttributeDefinition CustomSecurityAttributeDefinition) (*CustomSecurityAttributeDefinition, int, error) {
	var status int
	var newCustomSecurityAttributeDefinition CustomSecurityAttributeDefinition

	body, err := json.Marshal(customSecurityAttributeDefinition)
	if err != nil {
		return nil, status, fmt.Errorf("json.Marshal(): %v", err)
	}

	requestInput := PostHttpRequestInput{
		Body: body,
		OData: odata.Query{
			Metadata: odata.MetadataFull,
		},
		ValidStatusCodes: []int{http.StatusCreated},
		Uri: Uri{
			Entity: customSecurityAttributeDefinitionEntity,
		},
	}

	resp, status, _, err := c.BaseClient.Post(ctx, requestInput)
	if err != nil {
		return nil, status, fmt.Errorf("CustomSecurityAttributeDefinitionClient.BaseClient.Post(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	if err := json.Unmarshal(respBody, &newCustomSecurityAttributeDefinition); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal():%v", err)
	}

	return &newCustomSecurityAttributeDefinition, status, nil
}

// Get returns a single CustomSecurityAttributeDefinition, HTTP status code, and any errors
func (c *CustomSecurityAttributeDefinitionClient) Get(ctx context.Context, id string, query odata.Query) (*CustomSecurityAttributeDefinition, int, error) {
	var customSecurityAttributeDefinition CustomSecurityAttributeDefinition

	resp, status, _, err := c.BaseClient.Get(
		ctx,
		GetHttpRequestInput{
			ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
			OData:                  query,
			ValidStatusCodes:       []int{http.StatusOK},
			Uri: Uri{
				Entity: fmt.Sprintf("%s/%s", customSecurityAttributeDefinitionEntity, id),
			},
		},
	)
	if err != nil {
		return nil, status, fmt.Errorf("CustomSecurityAttributeDefinitionClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	if err := json.Unmarshal(respBody, &customSecurityAttributeDefinition); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &customSecurityAttributeDefinition, status, nil
}

// Update will update a single CustomSecurityAttributeDefinition entity returning the HTTP status
// code and any errors
func (c *CustomSecurityAttributeDefinitionClient) Update(ctx context.Context, customSecurityAttributeDefinition CustomSecurityAttributeDefinition) (int, error) {
	var status int

	if customSecurityAttributeDefinition.ID == nil {
		return status, fmt.Errorf("cannot update customSecurityAttributeDefinition with a nil ID")
	}

	id := *customSecurityAttributeDefinition.ID
	customSecurityAttributeDefinition.ID = nil

	body, err := json.Marshal(customSecurityAttributeDefinition)
	if err != nil {
		return status, fmt.Errorf("json.Marshal(): %v", err)
	}

	_, status, _, err = c.BaseClient.Patch(
		ctx,
		PatchHttpRequestInput{
			Body:                   body,
			ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
			ValidStatusCodes: []int{
				http.StatusOK,
				http.StatusNoContent,
			},
			Uri: Uri{
				Entity: fmt.Sprintf("%s/%s", customSecurityAttributeDefinitionEntity, id),
			},
		},
	)
	if err != nil {
		return status, fmt.Errorf("CustomSecurityAttributeDefinitionClient.BaseClient.Patch(): %v", err)
	}

	return status, nil
}

// Delete removes an instance of CustomSecurityAttributeDefinition by `id`
func (c *CustomSecurityAttributeDefinitionClient) Delete(ctx context.Context, id string) (int, error) {
	_, status, _, err := c.BaseClient.Delete(
		ctx,
		DeleteHttpRequestInput{
			ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
			ValidStatusCodes:       []int{http.StatusNoContent},
			Uri: Uri{
				Entity: fmt.Sprintf("%s/%s", customSecurityAttributeDefinitionEntity, id),
			},
		},
	)
	if err != nil {
		return status, fmt.Errorf("CustomSecurityAttributeDefinitionClient.BaseClient.Delete(): %v", err)
	}

	return status, nil
}

func (c *CustomSecurityAttributeDefinitionClient) Deactivate(ctx context.Context, id string) (int, error) {
	var status int
	var customSecurityAttributeDefinition CustomSecurityAttributeDefinition

	customSecurityAttributeDefinition.Status = utils.StringPtr("Deprecated")

	body, err := json.Marshal(customSecurityAttributeDefinition)
	if err != nil {
		return status, fmt.Errorf("json.Marshal(): %v", err)
	}

	_, status, _, err = c.BaseClient.Patch(
		ctx,
		PatchHttpRequestInput{
			Body:                   body,
			ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
			ValidStatusCodes: []int{
				http.StatusOK,
				http.StatusNoContent,
			},
			Uri: Uri{
				Entity: fmt.Sprintf("%s/%s", customSecurityAttributeDefinitionEntity, id),
			},
		},
	)
	if err != nil {
		return status, fmt.Errorf("customSecurityAttributeDefinitionClient.BaseClient.Patch(): %v", err)
	}

	return status, nil
}
