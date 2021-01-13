package clients

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"

	"github.com/manicminer/hamilton/base"
	"github.com/manicminer/hamilton/base/odata"
	"github.com/manicminer/hamilton/models"
)

// ServicePrincipalsClient performs operations on Service Principals.
type ServicePrincipalsClient struct {
	BaseClient base.Client
}

// NewServicePrincipalsClient returns a new ServicePrincipalsClient.
func NewServicePrincipalsClient(tenantId string) *ServicePrincipalsClient {
	return &ServicePrincipalsClient{
		BaseClient: base.NewClient(base.VersionBeta, tenantId),
	}
}

// List returns a list of Service Principals, optionally filtered using OData.
func (c *ServicePrincipalsClient) List(ctx context.Context, filter string) (*[]models.ServicePrincipal, int, error) {
	params := url.Values{}
	if filter != "" {
		params.Add("$filter", filter)
	}
	resp, status, _, err := c.BaseClient.Get(ctx, base.GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: base.Uri{
			Entity:      "/servicePrincipals",
			Params:      params,
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var data struct {
		ServicePrincipals []models.ServicePrincipal `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, err
	}
	return &data.ServicePrincipals, status, nil
}

// Create creates a new Service Principal.
func (c *ServicePrincipalsClient) Create(ctx context.Context, servicePrincipal models.ServicePrincipal) (*models.ServicePrincipal, int, error) {
	var status int
	body, err := json.Marshal(servicePrincipal)
	if err != nil {
		return nil, status, err
	}
	resp, status, _, err := c.BaseClient.Post(ctx, base.PostHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusCreated},
		Uri: base.Uri{
			Entity:      "/servicePrincipals",
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var newServicePrincipal models.ServicePrincipal
	if err := json.Unmarshal(respBody, &newServicePrincipal); err != nil {
		return nil, status, err
	}
	return &newServicePrincipal, status, nil
}

// Get retrieves a Service Principal.
func (c *ServicePrincipalsClient) Get(ctx context.Context, id string) (*models.ServicePrincipal, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, base.GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: base.Uri{
			Entity:      fmt.Sprintf("/servicePrincipals/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var servicePrincipal models.ServicePrincipal
	if err := json.Unmarshal(respBody, &servicePrincipal); err != nil {
		return nil, status, err
	}
	return &servicePrincipal, status, nil
}

// Update amends an existing Service Principal.
func (c *ServicePrincipalsClient) Update(ctx context.Context, servicePrincipal models.ServicePrincipal) (int, error) {
	var status int
	if servicePrincipal.ID == nil {
		return status, errors.New("cannot update service principal with nil ID")
	}
	body, err := json.Marshal(servicePrincipal)
	if err != nil {
		return status, err
	}
	_, status, _, err = c.BaseClient.Patch(ctx, base.PatchHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: base.Uri{
			Entity:      fmt.Sprintf("/servicePrincipals/%s", *servicePrincipal.ID),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, err
	}
	return status, nil
}

// Delete removes a Service Principal.
func (c *ServicePrincipalsClient) Delete(ctx context.Context, id string) (int, error) {
	_, status, _, err := c.BaseClient.Delete(ctx, base.DeleteHttpRequestInput{
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: base.Uri{
			Entity:      fmt.Sprintf("/servicePrincipals/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, err
	}
	return status, nil
}

// ListOwners retrieves the owners of the specified Service Principal.
// id is the object ID of the service principal.
func (c *ServicePrincipalsClient) ListOwners(ctx context.Context, id string) (*[]string, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, base.GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: base.Uri{
			Entity:      fmt.Sprintf("/servicePrincipals/%s/owners", id),
			Params:      url.Values{"$select": []string{"id"}},
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var data struct {
		Owners []struct {
			Type string `json:"@odata.type"`
			Id   string `json:"id"`
		} `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, err
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
	resp, status, _, err := c.BaseClient.Get(ctx, base.GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: base.Uri{
			Entity:      fmt.Sprintf("/servicePrincipals/%s/owners/%s/$ref", servicePrincipalId, ownerId),
			Params:      url.Values{"$select": []string{"id,url"}},
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var data struct {
		Context string `json:"@odata.context"`
		Type    string `json:"@odata.type"`
		Id      string `json:"id"`
		Url     string `json:"url"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, err
	}
	return &data.Id, status, nil
}

// AddOwners adds a new owner to a Service Principal.
// First populate the Owners field of the ServicePrincipal using the AppendOwner method of the model, then call this method.
func (c *ServicePrincipalsClient) AddOwners(ctx context.Context, servicePrincipal *models.ServicePrincipal) (int, error) {
	var status int
	if servicePrincipal.ID == nil {
		return status, errors.New("cannot update service principal with nil ID")
	}
	if servicePrincipal.Owners == nil {
		return status, errors.New("cannot update service principal with nil Owners")
	}
	for _, owner := range *servicePrincipal.Owners {
		data := struct {
			Owner string `json:"@odata.id"`
		}{
			Owner: owner,
		}
		body, err := json.Marshal(data)
		if err != nil {
			return status, err
		}
		_, status, _, err = c.BaseClient.Post(ctx, base.PostHttpRequestInput{
			Body:             body,
			ValidStatusCodes: []int{http.StatusNoContent},
			Uri: base.Uri{
				Entity:      fmt.Sprintf("/servicePrincipals/%s/owners/$ref", *servicePrincipal.ID),
				HasTenantId: true,
			},
		})
		if err != nil {
			return status, err
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
					re := regexp.MustCompile("One or more removed object references do not exist")
					if re.MatchString(o.Error.String()) {
						return true
					}
				}
			}
			return false
		}

		_, status, _, err := c.BaseClient.Delete(ctx, base.DeleteHttpRequestInput{
			ValidStatusCodes: []int{http.StatusNoContent},
			ValidStatusFunc:  checkOwnerGone,
			Uri: base.Uri{
				Entity:      fmt.Sprintf("/servicePrincipals/%s/owners/%s/$ref", servicePrincipalId, ownerId),
				HasTenantId: true,
			},
		})
		if err != nil {
			return status, err
		}
	}
	return status, nil
}

// ListGroupMemberships returns a list of Groups the Service Principal is member of, optionally filtered using OData.
func (c *ServicePrincipalsClient) ListGroupMemberships(ctx context.Context, id string, filter string) (*[]models.Group, int, error) {
	params := url.Values{}
	if filter != "" {
		params.Add("$filter", filter)
	}
	resp, status, _, err := c.BaseClient.Get(ctx, base.GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: base.Uri{
			Entity:      fmt.Sprintf("/servicePrincipals/%s/transitiveMemberOf", id),
			Params:      params,
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var data struct {
		Groups []models.Group `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, err
	}
	return &data.Groups, status, nil
}
