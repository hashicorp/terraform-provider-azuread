// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/me"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
	"github.com/hashicorp/go-azure-sdk/sdk/claims"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/terraform-provider-azuread/internal/common"

	administrativeunits "github.com/hashicorp/terraform-provider-azuread/internal/services/administrativeunits/client"
	applications "github.com/hashicorp/terraform-provider-azuread/internal/services/applications/client"
	approleassignments "github.com/hashicorp/terraform-provider-azuread/internal/services/approleassignments/client"
	conditionalaccess "github.com/hashicorp/terraform-provider-azuread/internal/services/conditionalaccess/client"
	directoryobjects "github.com/hashicorp/terraform-provider-azuread/internal/services/directoryobjects/client"
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
	DirectoryObjects    *directoryobjects.Client
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

	var err error

	if client.AdministrativeUnits, err = administrativeunits.NewClient(o); err != nil {
		return fmt.Errorf("building clients for AdministrativeUnits: %v", err)
	}
	if client.Applications, err = applications.NewClient(o); err != nil {
		return fmt.Errorf("building clients for Applications: %v", err)
	}
	if client.AppRoleAssignments, err = approleassignments.NewClient(o); err != nil {
		return fmt.Errorf("building clients for AppRoleAssignments: %v", err)
	}
	if client.ConditionalAccess, err = conditionalaccess.NewClient(o); err != nil {
		return fmt.Errorf("building clients for ConditionalAccess: %v", err)
	}
	if client.DirectoryObjects, err = directoryobjects.NewClient(o); err != nil {
		return fmt.Errorf("building clients for DirectoryObjects: %v", err)
	}
	if client.DirectoryRoles, err = directoryroles.NewClient(o); err != nil {
		return fmt.Errorf("building clients for DirectoryRoles: %v", err)
	}
	if client.Domains, err = domains.NewClient(o); err != nil {
		return fmt.Errorf("building clients for Domains: %v", err)
	}
	if client.Groups, err = groups.NewClient(o); err != nil {
		return fmt.Errorf("building clients for Groups: %v", err)
	}
	if client.IdentityGovernance, err = identitygovernance.NewClient(o); err != nil {
		return fmt.Errorf("building clients for IdentityGovernance: %v", err)
	}
	if client.Invitations, err = invitations.NewClient(o); err != nil {
		return fmt.Errorf("building clients for Invitations: %v", err)
	}
	if client.Policies, err = policies.NewClient(o); err != nil {
		return fmt.Errorf("building clients for Policies: %v", err)
	}
	if client.ServicePrincipals, err = serviceprincipals.NewClient(o); err != nil {
		return fmt.Errorf("building clients for ServicePrincipals: %v", err)
	}
	if client.Synchronization, err = synchronization.NewClient(o); err != nil {
		return fmt.Errorf("building clients for Synchronization: %v", err)
	}
	if client.UserFlows, err = userflows.NewClient(o); err != nil {
		return fmt.Errorf("building clients for UserFlows: %v", err)
	}
	if client.Users, err = users.NewClient(o); err != nil {
		return fmt.Errorf("building clients for Users: %v", err)
	}

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
		log.Printf("[DEBUG] AzureAD Provider could not marshal access token claims for log output")
	case claimsJson == nil:
		log.Printf("[DEBUG] AzureAD Provider access token claims was nil")
	default:
		log.Printf("[DEBUG] AzureAD Provider access token claims: %s", claimsJson)
	}

	// Missing object ID of token holder will break many things
	client.ObjectID = client.Claims.ObjectId
	if client.ObjectID == "" {
		if strings.Contains(strings.ToLower(client.Claims.Scopes), "openid") {
			log.Printf("[DEBUG] Querying Microsoft Graph to discover authenticated user principal object ID because the `oid` claim was missing from the access token")
			resp, err := client.Users.MeClient.GetMe(ctx, me.DefaultGetMeOperationOptions())
			if err != nil {
				return fmt.Errorf("attempting to discover object ID for authenticated user principal: %+v", err)
			}

			if resp.Model != nil {
				return fmt.Errorf("attempting to discover object ID for authenticated user principal: response was nil")
			}

			id := resp.Model.Id
			if id == nil {
				return fmt.Errorf("attempting to discover object ID for authenticated user principal: returned object ID was nil")
			}

			client.ObjectID = *id
		} else {
			log.Printf("[DEBUG] Querying Microsoft Graph to discover authenticated service principal object ID because the `oid` claim was missing from the access token")
			options := serviceprincipal.ListServicePrincipalsOperationOptions{
				Filter: pointer.To(fmt.Sprintf("appId eq '%s'", client.ClientID)),
			}
			resp, err := client.ServicePrincipals.ServicePrincipalClient.ListServicePrincipals(ctx, options)
			if err != nil {
				return fmt.Errorf("attempting to discover object ID for authenticated service principal: %+v", err)
			}

			if resp.Model != nil && len(*resp.Model) != 1 {
				respLen := "nil"
				if resp.Model != nil {
					respLen = strconv.Itoa(len(*resp.Model))
				}
				return fmt.Errorf("attempting to discover object ID for authenticated service principal: unexpected number of results, expected 1, received %s", respLen)
			}

			id := (*resp.Model)[0].Id
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
