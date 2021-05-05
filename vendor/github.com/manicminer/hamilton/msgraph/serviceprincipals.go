package msgraph

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/manicminer/hamilton/odata"
)

// ServicePrincipalsClient performs operations on Service Principals.
type ServicePrincipalsClient struct {
	BaseClient Client
}

// NewServicePrincipalsClient returns a new ServicePrincipalsClient.
func NewServicePrincipalsClient(tenantId string) *ServicePrincipalsClient {
	return &ServicePrincipalsClient{
		BaseClient: NewClient(VersionBeta, tenantId),
	}
}

// List returns a list of Service Principals, optionally filtered using OData.
func (c *ServicePrincipalsClient) List(ctx context.Context, filter string) (*[]ServicePrincipal, int, error) {
	params := url.Values{}
	if filter != "" {
		params.Add("$filter", filter)
	}
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      "/servicePrincipals",
			Params:      params,
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("ServicePrincipalsClient.BaseClient.Get(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("ioutil.ReadAll(): %v", err)
	}
	var data struct {
		ServicePrincipals []ServicePrincipal `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &data.ServicePrincipals, status, nil
}

// Create creates a new Service Principal.
func (c *ServicePrincipalsClient) Create(ctx context.Context, servicePrincipal ServicePrincipal) (*ServicePrincipal, int, error) {
	var status int
	body, err := json.Marshal(servicePrincipal)
	if err != nil {
		return nil, status, fmt.Errorf("json.Marshal(): %v", err)
	}
	resp, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusCreated},
		Uri: Uri{
			Entity:      "/servicePrincipals",
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("ServicePrincipalsClient.BaseClient.Post(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("ioutil.ReadAll(): %v", err)
	}
	var newServicePrincipal ServicePrincipal
	if err := json.Unmarshal(respBody, &newServicePrincipal); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &newServicePrincipal, status, nil
}

// Get retrieves a Service Principal.
func (c *ServicePrincipalsClient) Get(ctx context.Context, id string) (*ServicePrincipal, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/servicePrincipals/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("ServicePrincipalsClient.BaseClient.Get(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("ioutil.ReadAll(): %v", err)
	}
	var servicePrincipal ServicePrincipal
	if err := json.Unmarshal(respBody, &servicePrincipal); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &servicePrincipal, status, nil
}

// Update amends an existing Service Principal.
func (c *ServicePrincipalsClient) Update(ctx context.Context, servicePrincipal ServicePrincipal) (int, error) {
	var status int
	if servicePrincipal.ID == nil {
		return status, errors.New("cannot update service principal with nil ID")
	}
	body, err := json.Marshal(servicePrincipal)
	if err != nil {
		return status, fmt.Errorf("json.Marshal(): %v", err)
	}
	_, status, _, err = c.BaseClient.Patch(ctx, PatchHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: Uri{
			Entity:      fmt.Sprintf("/servicePrincipals/%s", *servicePrincipal.ID),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, fmt.Errorf("ServicePrincipalsClient.BaseClient.Patch(): %v", err)
	}
	return status, nil
}

// Delete removes a Service Principal.
func (c *ServicePrincipalsClient) Delete(ctx context.Context, id string) (int, error) {
	_, status, _, err := c.BaseClient.Delete(ctx, DeleteHttpRequestInput{
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: Uri{
			Entity:      fmt.Sprintf("/servicePrincipals/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, fmt.Errorf("ServicePrincipalsClient.BaseClient.Delete(): %v", err)
	}
	return status, nil
}

// ListOwners retrieves the owners of the specified Service Principal.
// id is the object ID of the service principal.
func (c *ServicePrincipalsClient) ListOwners(ctx context.Context, id string) (*[]string, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/servicePrincipals/%s/owners", id),
			Params:      url.Values{"$select": []string{"id"}},
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("ServicePrincipalsClient.BaseClient.Get(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("ioutil.ReadAll(): %v", err)
	}
	var data struct {
		Owners []struct {
			Type string `json:"@odata.type"`
			Id   string `json:"id"`
		} `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	ret := make([]string, len(data.Owners))
	for i, v := range data.Owners {
		ret[i] = v.Id
	}
	return &ret, status, nil
}

// GetOwner retrieves a single owner for the specified Service Principal.
// servicePrincipalId is the object ID of the service principal.
// ownerId is the object ID of the owning object.
func (c *ServicePrincipalsClient) GetOwner(ctx context.Context, servicePrincipalId, ownerId string) (*string, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/servicePrincipals/%s/owners/%s/$ref", servicePrincipalId, ownerId),
			Params:      url.Values{"$select": []string{"id,url"}},
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("ServicePrincipalsClient.BaseClient.Get(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("ioutil.ReadAll(): %v", err)
	}
	var data struct {
		Context string `json:"@odata.context"`
		Type    string `json:"@odata.type"`
		Id      string `json:"id"`
		Url     string `json:"url"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &data.Id, status, nil
}

// AddOwners adds a new owner to a Service Principal.
// First populate the Owners field of the ServicePrincipal using the AppendOwner method of the model, then call this method.
func (c *ServicePrincipalsClient) AddOwners(ctx context.Context, servicePrincipal *ServicePrincipal) (int, error) {
	var status int
	if servicePrincipal.ID == nil {
		return status, errors.New("cannot update service principal with nil ID")
	}
	if servicePrincipal.Owners == nil {
		return status, errors.New("cannot update service principal with nil Owners")
	}
	for _, owner := range *servicePrincipal.Owners {
		// don't fail if an owner already exists
		checkOwnerAlreadyExists := func(resp *http.Response, o *odata.OData) bool {
			if resp.StatusCode == http.StatusBadRequest {
				if o.Error != nil {
					if o.Error.Match(odata.ErrorAddedObjectReferencesAlreadyExist) {
						return true
					}
				}
			}
			return false
		}

		data := struct {
			Owner string `json:"@odata.id"`
		}{
			Owner: owner,
		}
		body, err := json.Marshal(data)
		if err != nil {
			return status, fmt.Errorf("json.Marshal(): %v", err)
		}
		_, status, _, err = c.BaseClient.Post(ctx, PostHttpRequestInput{
			Body:             body,
			ValidStatusCodes: []int{http.StatusNoContent},
			ValidStatusFunc:  checkOwnerAlreadyExists,
			Uri: Uri{
				Entity:      fmt.Sprintf("/servicePrincipals/%s/owners/$ref", *servicePrincipal.ID),
				HasTenantId: true,
			},
		})
		if err != nil {
			return status, fmt.Errorf("ServicePrincipalsClient.BaseClient.Post(): %v", err)
		}
	}
	return status, nil
}

// RemoveOwners removes owners from a Service Principal.
// servicePrincipalId is the object ID of the service principal.
// ownerIds is a *[]string containing object IDs of owners to remove.
func (c *ServicePrincipalsClient) RemoveOwners(ctx context.Context, servicePrincipalId string, ownerIds *[]string) (int, error) {
	var status int
	if ownerIds == nil {
		return status, errors.New("cannot remove, nil ownerIds")
	}
	for _, ownerId := range *ownerIds {
		// check for ownership before attempting deletion
		if _, status, err := c.GetOwner(ctx, servicePrincipalId, ownerId); err != nil {
			if status == http.StatusNotFound {
				continue
			}
			return status, err
		}

		// despite the above check, sometimes owners are just gone
		checkOwnerGone := func(resp *http.Response, o *odata.OData) bool {
			if resp.StatusCode == http.StatusBadRequest {
				if o.Error != nil {
					if o.Error.Match(odata.ErrorRemovedObjectReferencesDoNotExist) {
						return true
					}
				}
			}
			return false
		}

		_, status, _, err := c.BaseClient.Delete(ctx, DeleteHttpRequestInput{
			ValidStatusCodes: []int{http.StatusNoContent},
			ValidStatusFunc:  checkOwnerGone,
			Uri: Uri{
				Entity:      fmt.Sprintf("/servicePrincipals/%s/owners/%s/$ref", servicePrincipalId, ownerId),
				HasTenantId: true,
			},
		})
		if err != nil {
			return status, fmt.Errorf("ServicePrincipalsClient.BaseClient.Delete(): %v", err)
		}
	}
	return status, nil
}

// ListGroupMemberships returns a list of Groups the Service Principal is member of, optionally filtered using OData.
func (c *ServicePrincipalsClient) ListGroupMemberships(ctx context.Context, id string, filter string) (*[]Group, int, error) {
	params := url.Values{}
	if filter != "" {
		params.Add("$filter", filter)
	}
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/servicePrincipals/%s/transitiveMemberOf", id),
			Params:      params,
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("ServicePrincipalsClient.BaseClient.Get(): %v", err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("ioutil.ReadAll(): %v", err)
	}
	var data struct {
		Groups []Group `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}
	return &data.Groups, status, nil
}
