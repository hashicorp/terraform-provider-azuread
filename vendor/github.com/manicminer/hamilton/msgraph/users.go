package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// UsersClient performs operations on Users.
type UsersClient struct {
	BaseClient Client
}

// NewUsersClient returns a new UsersClient.
func NewUsersClient(tenantId string) *UsersClient {
	return &UsersClient{
		BaseClient: NewClient(VersionBeta, tenantId),
	}
}

// List returns a list of Users, optionally filtered using OData.
func (c *UsersClient) List(ctx context.Context, filter string) (*[]User, int, error) {
	params := url.Values{}
	if filter != "" {
		params.Add("$filter", filter)
	}
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      "/users",
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
		Users []User `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, err
	}
	return &data.Users, status, nil
}

// Create creates a new User.
func (c *UsersClient) Create(ctx context.Context, user User) (*User, int, error) {
	var status int
	body, err := json.Marshal(user)
	if err != nil {
		return nil, status, err
	}
	resp, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusCreated},
		Uri: Uri{
			Entity:      "/users",
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var newUser User
	if err := json.Unmarshal(respBody, &newUser); err != nil {
		return nil, status, err
	}
	return &newUser, status, nil
}

// Get retrieves a User.
func (c *UsersClient) Get(ctx context.Context, id string) (*User, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/users/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var user User
	if err := json.Unmarshal(respBody, &user); err != nil {
		return nil, status, err
	}
	return &user, status, nil
}

// Update amends an existing User.
func (c *UsersClient) Update(ctx context.Context, user User) (int, error) {
	var status int
	body, err := json.Marshal(user)
	if err != nil {
		return status, err
	}
	_, status, _, err = c.BaseClient.Patch(ctx, PatchHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: Uri{
			Entity:      fmt.Sprintf("/users/%s", *user.ID),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, err
	}
	return status, nil
}

// Delete removes a User.
func (c *UsersClient) Delete(ctx context.Context, id string) (int, error) {
	_, status, _, err := c.BaseClient.Delete(ctx, DeleteHttpRequestInput{
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: Uri{
			Entity:      fmt.Sprintf("/users/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, err
	}
	return status, nil
}

// ListGroupMemberships returns a list of Groups the user is member of, optionally filtered using OData.
func (c *UsersClient) ListGroupMemberships(ctx context.Context, id string, filter string) (*[]Group, int, error) {
	params := url.Values{}
	if filter != "" {
		params.Add("$filter", filter)
	}
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/users/%s/transitiveMemberOf", id),
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
		Groups []Group `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, err
	}
	return &data.Groups, status, nil
}
