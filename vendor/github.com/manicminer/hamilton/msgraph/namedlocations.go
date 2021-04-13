package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/manicminer/hamilton/internal/utils"
	"github.com/manicminer/hamilton/odata"
)

// NamedLocationsClient performs operations on Named Locations.
type NamedLocationsClient struct {
	BaseClient Client
}

// NewNamedLocationsClient returns a new NamedLocationsClient.
func NewNamedLocationsClient(tenantId string) *NamedLocationsClient {
	return &NamedLocationsClient{
		BaseClient: NewClient(Version10, tenantId),
	}
}

// List returns a list of Named Locations, optionally filtered using OData.
func (c *NamedLocationsClient) List(ctx context.Context, filter string) (*[]NamedLocation, int, error) {
	params := url.Values{}
	if filter != "" {
		params.Add("$filter", filter)
	}

	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      "/identity/conditionalAccess/namedLocations",
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
		NamedLocations *[]json.RawMessage `json:"value"`
	}

	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, err
	}

	if data.NamedLocations == nil {
		return nil, status, err
	}
	// The Graph API returns a mixture of types, this loop matches up the result to the appropriate model
	var ret []NamedLocation
	for _, namedLocation := range *data.NamedLocations {
		var o odata.OData
		if err := json.Unmarshal(namedLocation, &o); err != nil {
			return nil, status, err
		}

		if o.Type == nil {
			continue
		}
		switch *o.Type {
		case "#microsoft.graph.countryNamedLocation":
			var loc CountryNamedLocation
			if err := json.Unmarshal(namedLocation, &loc); err != nil {
				return nil, status, err
			}
			ret = append(ret, loc)
		case "#microsoft.graph.ipNamedLocation":
			var loc IPNamedLocation
			if err := json.Unmarshal(namedLocation, &loc); err != nil {
				return nil, status, err
			}
			ret = append(ret, loc)
		}
	}

	return &ret, status, nil

}

// Delete removes a Named Location.
func (c *NamedLocationsClient) Delete(ctx context.Context, id string) (int, error) {
	_, status, _, err := c.BaseClient.Delete(ctx, DeleteHttpRequestInput{
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: Uri{
			Entity:      fmt.Sprintf("/identity/conditionalAccess/namedLocations/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, err
	}
	return status, nil
}

// CreateIP creates a new IP Named Location.
func (c *NamedLocationsClient) CreateIP(ctx context.Context, ipNamedLocation IPNamedLocation) (*IPNamedLocation, int, error) {
	var status int

	ipNamedLocation.ODataType = utils.StringPtr("#microsoft.graph.ipNamedLocation")
	body, err := json.Marshal(ipNamedLocation)
	if err != nil {
		return nil, status, err
	}
	resp, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusCreated},
		Uri: Uri{
			Entity:      "/identity/conditionalAccess/namedLocations",
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var newIPNamedLocation IPNamedLocation
	if err := json.Unmarshal(respBody, &newIPNamedLocation); err != nil {
		return nil, status, err
	}
	return &newIPNamedLocation, status, nil
}

// CreateCountry creates a new Country Named Location.
func (c *NamedLocationsClient) CreateCountry(ctx context.Context, countryNamedLocation CountryNamedLocation) (*CountryNamedLocation, int, error) {
	var status int

	countryNamedLocation.ODataType = utils.StringPtr("#microsoft.graph.countryNamedLocation")

	body, err := json.Marshal(countryNamedLocation)
	if err != nil {
		return nil, status, err
	}
	resp, status, _, err := c.BaseClient.Post(ctx, PostHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusCreated},
		Uri: Uri{
			Entity:      "/identity/conditionalAccess/namedLocations",
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var newCountryNamedLocation CountryNamedLocation
	if err := json.Unmarshal(respBody, &newCountryNamedLocation); err != nil {
		return nil, status, err
	}
	return &newCountryNamedLocation, status, nil
}

// GetIP retrieves an IP Named Location.
func (c *NamedLocationsClient) GetIP(ctx context.Context, id string) (*IPNamedLocation, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/identity/conditionalAccess/namedLocations/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var ipNamedLocation IPNamedLocation
	if err := json.Unmarshal(respBody, &ipNamedLocation); err != nil {
		return nil, status, err
	}
	return &ipNamedLocation, status, nil
}

// GetCountry retrieves an Country Named Location.
func (c *NamedLocationsClient) GetCountry(ctx context.Context, id string) (*CountryNamedLocation, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity:      fmt.Sprintf("/identity/conditionalAccess/namedLocations/%s", id),
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, status, err
	}
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	var countryNamedLocation CountryNamedLocation
	if err := json.Unmarshal(respBody, &countryNamedLocation); err != nil {
		return nil, status, err
	}
	return &countryNamedLocation, status, nil
}

// UpdateIP amends an existing IP Named Location.
func (c *NamedLocationsClient) UpdateIP(ctx context.Context, ipNamedLocation IPNamedLocation) (int, error) {
	var status int

	body, err := json.Marshal(ipNamedLocation)
	if err != nil {
		return status, err
	}
	_, status, _, err = c.BaseClient.Patch(ctx, PatchHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: Uri{
			Entity:      fmt.Sprintf("/identity/conditionalAccess/namedLocations/%s", *ipNamedLocation.ID),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, err
	}
	return status, nil
}

// UpdateCountry amends an existing Country Named Location.
func (c *NamedLocationsClient) UpdateCountry(ctx context.Context, countryNamedLocation CountryNamedLocation) (int, error) {
	var status int

	body, err := json.Marshal(countryNamedLocation)
	if err != nil {
		return status, err
	}
	_, status, _, err = c.BaseClient.Patch(ctx, PatchHttpRequestInput{
		Body:             body,
		ValidStatusCodes: []int{http.StatusNoContent},
		Uri: Uri{
			Entity:      fmt.Sprintf("/identity/conditionalAccess/namedLocations/%s", *countryNamedLocation.ID),
			HasTenantId: true,
		},
	})
	if err != nil {
		return status, err
	}
	return status, nil
}
