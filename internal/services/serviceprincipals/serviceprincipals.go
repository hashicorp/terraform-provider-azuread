package serviceprincipals

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/utils"
	"github.com/manicminer/hamilton/msgraph"
)

func expandSamlSingleSignOn(in []interface{}) *msgraph.SamlSingleSignOnSettings {
	result := msgraph.SamlSingleSignOnSettings{}
	if len(in) == 0 || in[0] == nil {
		return &result
	}

	samlSingleSignOnSettings := in[0].(map[string]interface{})

	result.RelayState = utils.String(samlSingleSignOnSettings["relay_state"].(string))

	return &result
}

func flattenSamlSingleSignOn(in *msgraph.SamlSingleSignOnSettings) []map[string]interface{} {
	if in == nil {
		return []map[string]interface{}{}
	}

	relayState := ""
	if in.RelayState != nil {
		relayState = *in.RelayState
	}

	return []map[string]interface{}{{
		"relay_state": relayState,
	}}
}

func findByAppId(ctx context.Context, client *msgraph.ServicePrincipalsClient, appId string) (*msgraph.ServicePrincipal, error) {
	var servicePrincipal *msgraph.ServicePrincipal

	result, _, err := client.List(ctx, odata.Query{Filter: fmt.Sprintf("appId eq '%s'", appId)})
	if err != nil {
		return nil, fmt.Errorf("could not list existing service principals")
	}
	if result != nil {
		for _, r := range *result {
			if r.AppId != nil && strings.EqualFold(*r.AppId, appId) {
				servicePrincipal = &r
				break
			}
		}
	}

	return servicePrincipal, nil
}

func findByAppIdWithTimeout(ctx context.Context, timeout time.Duration, client *msgraph.ServicePrincipalsClient, appId string) (*msgraph.ServicePrincipal, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, timeout)
	defer cancel()

	unmarshal := func(resp *http.Response) (*msgraph.ServicePrincipal, error) {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("listing service principals: %v", err)
		}

		// Close and reassign the response body
		resp.Body.Close()
		resp.Body = io.NopCloser(bytes.NewBuffer(respBody))

		var data struct {
			ServicePrincipals []msgraph.ServicePrincipal `json:"value"`
		}
		if err := json.Unmarshal(respBody, &data); err != nil {
			return nil, fmt.Errorf("unmarshaling service principals: %v", err)
		}

		if len(data.ServicePrincipals) == 0 {
			return nil, nil
		} else if len(data.ServicePrincipals) > 1 {
			return nil, fmt.Errorf("unexpected number of results, should have received 0 or 1, got %d", len(data.ServicePrincipals))
		}

		if data.ServicePrincipals[0].AppId == nil || !strings.EqualFold(*data.ServicePrincipals[0].AppId, appId) {
			return nil, fmt.Errorf("returned service principal did not have a matching appId, expected %q, received %q", appId, *data.ServicePrincipals[0].AppId)
		}

		return &data.ServicePrincipals[0], nil
	}

	notReplicated := func(resp *http.Response, o *odata.OData) bool {
		sp, err := unmarshal(resp)
		if err == nil && sp == nil {
			return false
		}
		return false
	}

	resp, _, _, err := client.BaseClient.Get(ctx, msgraph.GetHttpRequestInput{
		ConsistencyFailureFunc: notReplicated,
		DisablePaging:          true,
		OData:                  odata.Query{Filter: fmt.Sprintf("appId eq '%s'", appId)},
		ValidStatusCodes:       []int{http.StatusOK},
		Uri: msgraph.Uri{
			Entity:      "/servicePrincipals",
			HasTenantId: true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("listing service principals: %v", err)
	}

	return unmarshal(resp)
}
