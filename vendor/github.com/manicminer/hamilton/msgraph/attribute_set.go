package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

const (
	attributeSetEntity = "/directory/attributeSets"
)

type AttributeSetClient struct {
	BaseClient Client
}

func NewAttributeSetClient() *AttributeSetClient {
	return &AttributeSetClient{
		BaseClient: NewClient(Version10),
	}
}

func (c *AttributeSetClient) List(ctx context.Context, query odata.Query) (*[]AttributeSet, int, error) {
	resp, status, _, err := c.BaseClient.Get(
		ctx,
		GetHttpRequestInput{
			OData:            query,
			ValidStatusCodes: []int{http.StatusOK},
			Uri: Uri{
				Entity: attributeSetEntity,
			},
		},
	)
	if err != nil {
		return nil, status, fmt.Errorf("AttributeSet.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		AttributeSets []AttributeSet `json:"value"`
	}

	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &data.AttributeSets, status, nil
}

func (c *AttributeSetClient) Create(ctx context.Context, attributeSet AttributeSet) (*AttributeSet, int, error) {
	var status int
	var newAttributeSet AttributeSet

	body, err := json.Marshal(attributeSet)
	if err != nil {
		return nil, status, fmt.Errorf("json.Marshal(): %v", err)
	}

	requestInput := PostHttpRequestInput{
		Body: body,
		OData: odata.Query{
			Metadata: odata.MetadataFull,
		},
		ValidStatusCodes: []int{
			http.StatusCreated,
			http.StatusOK,
		},
		Uri: Uri{
			Entity: attributeSetEntity,
		},
	}

	resp, status, _, err := c.BaseClient.Post(ctx, requestInput)
	if err != nil {
		return nil, status, fmt.Errorf("AttributeSetClient.BaseClient.Post(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	if err := json.Unmarshal(respBody, &newAttributeSet); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal():%v", err)
	}

	return &newAttributeSet, status, nil
}

func (c *AttributeSetClient) Get(ctx context.Context, id string, query odata.Query) (*AttributeSet, int, error) {
	var AttributeSet AttributeSet

	resp, status, _, err := c.BaseClient.Get(
		ctx,
		GetHttpRequestInput{
			ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
			OData:                  query,
			ValidStatusCodes:       []int{http.StatusOK},
			Uri: Uri{
				Entity: fmt.Sprintf("%s/%s", attributeSetEntity, id),
			},
		},
	)
	if err != nil {
		return nil, status, fmt.Errorf("AttributeSetClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	if err := json.Unmarshal(respBody, &AttributeSet); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	return &AttributeSet, status, nil
}

func (c *AttributeSetClient) Update(ctx context.Context, AttributeSet AttributeSet) (int, error) {
	var status int

	if AttributeSet.ID == nil {
		return status, fmt.Errorf("cannot update AttributeSet with a nil ID")
	}

	id := *AttributeSet.ID
	AttributeSet.ID = nil

	body, err := json.Marshal(AttributeSet)
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
				Entity: fmt.Sprintf("%s/%s", attributeSetEntity, id),
			},
		},
	)
	if err != nil {
		return status, fmt.Errorf("AttributeSetClient.BaseClient.Patch(): %v", err)
	}

	return status, nil
}
