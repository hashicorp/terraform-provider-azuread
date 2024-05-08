// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/claims"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/terraform-provider-azuread/internal/common"

	administrativeunits "github.com/hashicorp/terraform-provider-azuread/internal/services/administrativeunits/client"
	applications "github.com/hashicorp/terraform-provider-azuread/internal/services/applications/client"
	approleassignments "github.com/hashicorp/terraform-provider-azuread/internal/services/approleassignments/client"
	conditionalaccess "github.com/hashicorp/terraform-provider-azuread/internal/services/conditionalaccess/client"
	directoryroles "github.com/hashicorp/terraform-provider-azuread/internal/services/directoryroles/client"
	domains "github.com/hashicorp/terraform-provider-azuread/internal/services/domains/client"
	groups "github.com/hashicorp/terraform-provider-azuread/internal/services/groups/client"
	identitygovernance "github.com/hashicorp/terraform-provider-azuread/internal/services/identitygovernance/client"
	invitations "github.com/hashicorp/terraform-provider-azuread/internal/services/invitations/client"
	policies "github.com/hashicorp/terraform-provider-azuread/internal/services/policies/client"
	serviceprincipals "github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/client"
	synchronization "github.com/hashicorp/terraform-provider-azuread/internal/services/synchronization/client"
	userflows "github.com/hashicorp/terraform-provider-azuread/internal/services/userflows/client"
	users "github.com/hashicorp/terraform-provider-azuread/internal/services/users/client"
)

// Client contains the handles to all the specific Azure AD resource classes' respective clients
type Client struct {
	Environment environments.Environment
	TenantID    string
	ClientID    string
	ObjectID    string
	Claims      *claims.Claims

	TerraformVersion string

	StopContext context.Context

	AdministrativeUnits *administrativeunits.Client
	Applications        *applications.Client
	AppRoleAssignments  *approleassignments.Client
	ConditionalAccess   *conditionalaccess.Client
	DirectoryRoles      *directoryroles.Client
	Domains             *domains.Client
	Groups              *groups.Client
	IdentityGovernance  *identitygovernance.Client
	Invitations         *invitations.Client
	Policies            *policies.Client
	ServicePrincipals   *serviceprincipals.Client
	Synchronization     *synchronization.Client
	UserFlows           *userflows.Client
	Users               *users.Client
}

func (client *Client) build(ctx context.Context, o *common.ClientOptions) error {
	client.StopContext = ctx

	client.AdministrativeUnits = administrativeunits.NewClient(o)
	client.Applications = applications.NewClient(o)
	client.AppRoleAssignments = approleassignments.NewClient(o)
	client.Domains = domains.NewClient(o)
	client.ConditionalAccess = conditionalaccess.NewClient(o)
	client.DirectoryRoles = directoryroles.NewClient(o)
	client.Groups = groups.NewClient(o)
	client.IdentityGovernance = identitygovernance.NewClient(o)
	client.Invitations = invitations.NewClient(o)
	client.Policies = policies.NewClient(o)
	client.ServicePrincipals = serviceprincipals.NewClient(o)
	client.Synchronization = synchronization.NewClient(o)
	client.UserFlows = userflows.NewClient(o)
	client.Users = users.NewClient(o)

	// Acquire an access token upfront, so we can decode the JWT and populate the claims
	token, err := o.Authorizer.Token(ctx, &http.Request{})
	if err != nil {
		return fmt.Errorf("unable to obtain access token: %v", err)
	}

	client.Claims, err = claims.ParseClaims(token)
	if err != nil {
		return fmt.Errorf("unable to parse claims in access token: %v", err)
	}

	// Log the claims for debugging
	claimsJson, err := json.Marshal(client.Claims)
	switch {
	case err != nil:
		log.Printf("[DEBUG] AzureAD Provider could not marshal access token claims for log outout")
	case claimsJson == nil:
		log.Printf("[DEBUG] AzureAD Provider marshaled access token claims was nil")
	default:
		log.Printf("[DEBUG] AzureAD Provider access token claims: %s", claimsJson)
	}

	// Missing object ID of token holder will break many things
	client.ObjectID = client.Claims.ObjectId
	if client.ObjectID == "" {
		if strings.Contains(strings.ToLower(client.Claims.Scopes), "openid") {
			log.Printf("[DEBUG] Querying Microsoft Graph to discover authenticated user principal object ID because the `oid` claim was missing from the access token")
			result, _, err := client.Users.MeClient.Get(ctx, odata.Query{})
			if err != nil {
				return fmt.Errorf("attempting to discover object ID for authenticated user principal: %+v", err)
			}

			id := result.ID
			if id == nil {
				return fmt.Errorf("attempting to discover object ID for authenticated user principal: returned object ID was nil")
			}

			client.ObjectID = *id
		} else {
			log.Printf("[DEBUG] Querying Microsoft Graph to discover authenticated service principal object ID because the `oid` claim was missing from the access token")
			query := odata.Query{
				Filter: fmt.Sprintf("appId eq '%s'", client.ClientID),
			}
			result, _, err := client.ServicePrincipals.ServicePrincipalsClient.List(ctx, query)
			if err != nil {
				return fmt.Errorf("attempting to discover object ID for authenticated service principal: %+v", err)
			}

			if len(*result) != 1 {
				return fmt.Errorf("attempting to discover object ID for authenticated service principal: unexpected number of results, expected 1, received %d", len(*result))
			}

			id := (*result)[0].ID()
			if id == nil {
				return fmt.Errorf("attempting to discover object ID for authenticated service principal: returned object ID was nil")
			}

			client.ObjectID = *id
		}
	}

	if client.ObjectID == "" {
		return fmt.Errorf("parsing claims in access token: oid claim is empty")
	}

	return nil
}
