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

// ApplicationsClient performs operations on Applications.
type ApplicationsClient struct {
	BaseClient base.Client
}

// NewApplicationsClient returns a new ApplicationsClient
func NewApplicationsClient(tenantId string) *ApplicationsClient {
	return &ApplicationsClient{
		BaseClient: base.NewClient(base.VersionBeta, tenantId),
	}
}

// List returns a list of Applications, optionally filtered using OData.
func (c *ApplicationsClient) List(ctx context.Context, filter string) (*[]models.Application, int, error) {
	params := url.Values{}
	if filter != "" {
		params.Add("$filter", filter)
	}
	resp, status, _, err := c.BaseClient.Get(ctx, base.GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: base.Uri{
			Entity:      "/applications",
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
		Applications []models.Application `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, err
	}
	return &data.Applications, status, nil
}

// Create creates a new Application.
func (c *ApplicationsClient) Create(ctx context.Context, application models.Application) (*models.Application, int, error) {
	var status int
	body, err := json.Marshal(application)
	if err != nil {
		return nil, status, err
	}
	resp, status, _, err := c.BaseClient.Post(ctx, base.PostHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusCreated},
		Uri: base.Uri{
			Entity:      "/applications",
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var newApplication models.Application
	if err := json.Unmarshal(respBody, &newApplication); err != nil {
		return nil, status, err
	}
	return &newApplication, status, nil
}

// Get retrieves an Application manifest.
func (c *ApplicationsClient) Get(ctx context.Context, id string) (*models.Application, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, base.GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: base.Uri{
			Entity:      fmt.Sprintf("/applications/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var application models.Application
	if err := json.Unmarshal(respBody, &application); err != nil {
		return nil, status, err
	}
	return &application, status, nil
}

// Update amends the manifest of an existing Application.
func (c *ApplicationsClient) Update(ctx context.Context, application models.Application) (int, error) {
	var status int
	if application.ID == nil {
		return status, errors.New("cannot update application with nil ID")
	}
	body, err := json.Marshal(application)
	if err != nil {
		return status, err
	}
	_, status, _, err = c.BaseClient.Patch(ctx, base.PatchHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: base.Uri{
			Entity:      fmt.Sprintf("/applications/%s", *application.ID),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, err
	}
	return status, nil
}

// Delete removes an Application.
func (c *ApplicationsClient) Delete(ctx context.Context, id string) (int, error) {
	_, status, _, err := c.BaseClient.Delete(ctx, base.DeleteHttpRequestInput{
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: base.Uri{
			Entity:      fmt.Sprintf("/applications/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, err
	}
	return status, nil
}

// AddKey appends a new key credential to an Application.
func (c *ApplicationsClient) AddKey(ctx context.Context, applicationId string, keyCredential models.KeyCredential) (*models.KeyCredential, int, error) {
	var status int
	body, err := json.Marshal(keyCredential)
	if err != nil {
		return nil, status, err
	}
	resp, status, _, err := c.BaseClient.Post(ctx, base.PostHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusOK, http.StatusCreated},
		Uri: base.Uri{
			Entity:      fmt.Sprintf("/applications/%s/addKey", applicationId),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var newKeyCredential models.KeyCredential
	if err := json.Unmarshal(respBody, &newKeyCredential); err != nil {
		return nil, status, err
	}
	return &newKeyCredential, status, nil
}

// ListOwners retrieves the owners of the specified Application.
// id is the object ID of the application.
func (c *ApplicationsClient) ListOwners(ctx context.Context, id string) (*[]string, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, base.GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: base.Uri{
			Entity:      fmt.Sprintf("/applications/%s/owners", id),
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

// GetOwner retrieves a single owner for the specified Application.
// applicationId is the object ID of the application.
// ownerId is the object ID of the owning object.
func (c *ApplicationsClient) GetOwner(ctx context.Context, applicationId, ownerId string) (*string, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, base.GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: base.Uri{
			Entity:      fmt.Sprintf("/applications/%s/owners/%s/$ref", applicationId, ownerId),
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

// AddOwners adds a new owner to an Application.
// First populate the Owners field of the Application using the AppendOwner method of the model, then call this method.
func (c *ApplicationsClient) AddOwners(ctx context.Context, application *models.Application) (int, error) {
	var status int
	if application.ID == nil {
		return status, errors.New("cannot update application with nil ID")
	}
	if application.Owners == nil {
		return status, errors.New("cannot update application with nil Owners")
	}
	for _, owner := range *application.Owners {
		// don't fail if an owner already exists
		checkOwnerAlreadyExists := func(resp *http.Response, o *odata.OData) bool {
			if resp.StatusCode == http.StatusBadRequest {
				if o.Error != nil {
					re := regexp.MustCompile("One or more added object references already exist")
					if re.MatchString(o.Error.String()) {
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
			return status, err
		}
		_, status, _, err = c.BaseClient.Post(ctx, base.PostHttpRequestInput{
			Body:             body,
			ValidStatusCodes: []int{http.StatusNoContent},
			ValidStatusFunc:  checkOwnerAlreadyExists,
			Uri: base.Uri{
				Entity:      fmt.Sprintf("/applications/%s/owners/$ref", *application.ID),
				HasTenantId: true,
			},
		})
		if err != nil {
			return status, err
		}
	}
	return status, nil
}

// RemoveOwners removes owners from an Application.
// applicationId is the object ID of the application.
// ownerIds is a *[]string containing object IDs of owners to remove.
func (c *ApplicationsClient) RemoveOwners(ctx context.Context, applicationId string, ownerIds *[]string) (int, error) {
	var status int
	if ownerIds == nil {
		return status, errors.New("cannot remove, nil ownerIds")
	}
	for _, ownerId := range *ownerIds {
		// check for ownership before attempting deletion
		if _, status, err := c.GetOwner(ctx, applicationId, ownerId); err != nil {
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

		var err error
		_, status, _, err = c.BaseClient.Delete(ctx, base.DeleteHttpRequestInput{
			ValidStatusCodes: []int{http.StatusNoContent},
			ValidStatusFunc:  checkOwnerGone,
			Uri: base.Uri{
				Entity:      fmt.Sprintf("/applications/%s/owners/%s/$ref", applicationId, ownerId),
				HasTenantId: true,
			},
		})
		if err != nil {
			return status, err
		}
	}
	return status, nil
}
