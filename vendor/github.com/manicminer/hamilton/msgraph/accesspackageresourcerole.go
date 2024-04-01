package msgraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type AccessPackageResourceRoleClient struct {
	BaseClient Client
}

func NewAccessPackageResourceRoleClient() *AccessPackageResourceRoleClient {
	return &AccessPackageResourceRoleClient{
		BaseClient: NewClient(VersionBeta),
	}
}

// List retrieves a list of AccessPackageResourceRoles for a specific accessPackageResource for a particular catalog / originSystem
// This method requires us to use an Odata Filter / Expand to function correctly
func (c *AccessPackageResourceRoleClient) List(ctx context.Context, catalogId string, originSystem AccessPackageResourceOriginSystem, accessPackageResourceId string) (*[]AccessPackageResourceRole, int, error) {
	resp, status, _, err := c.BaseClient.Get(ctx, GetHttpRequestInput{
		ConsistencyFailureFunc: RetryOn404ConsistencyFailureFunc,
		OData: odata.Query{
			Filter: fmt.Sprintf("originSystem eq '%s' and accessPackageResource/id eq '%s'", originSystem, accessPackageResourceId),
			Expand: odata.Expand{
				Relationship: "accessPackageResource",
			},
		},
		ValidStatusCodes: []int{http.StatusOK},
		Uri: Uri{
			Entity: fmt.Sprintf("/identityGovernance/entitlementManagement/accessPackageCatalogs/%s/accessPackageResourceRoles", catalogId),
		},
	})
	if err != nil {
		return nil, status, fmt.Errorf("AccessPackageResourceRoleClient.BaseClient.Get(): %v", err)
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, status, fmt.Errorf("io.ReadAll(): %v", err)
	}

	var data struct {
		AccessPackageResourceRoles []AccessPackageResourceRole `json:"value"`
	}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return nil, status, fmt.Errorf("json.Unmarshal(): %v", err)
	}

	AccessPackageResourceRoles := data.AccessPackageResourceRoles

	if len(AccessPackageResourceRoles) == 0 {
		return nil, http.StatusNotFound, fmt.Errorf("no AccessPackageResourceRoles found with catalogId %v, originSystem %v and accessPackageResourceId %v", catalogId, originSystem, accessPackageResourceId)
	}

	return &AccessPackageResourceRoles, status, nil
}
