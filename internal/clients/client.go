package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/manicminer/hamilton/auth"
	"github.com/manicminer/hamilton/environments"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	administrativeunits "github.com/hashicorp/terraform-provider-azuread/internal/services/administrativeunits/client"
	applications "github.com/hashicorp/terraform-provider-azuread/internal/services/applications/client"
	approleassignments "github.com/hashicorp/terraform-provider-azuread/internal/services/approleassignments/client"
	conditionalaccess "github.com/hashicorp/terraform-provider-azuread/internal/services/conditionalaccess/client"
	directoryroles "github.com/hashicorp/terraform-provider-azuread/internal/services/directoryroles/client"
	domains "github.com/hashicorp/terraform-provider-azuread/internal/services/domains/client"
	groups "github.com/hashicorp/terraform-provider-azuread/internal/services/groups/client"
	invitations "github.com/hashicorp/terraform-provider-azuread/internal/services/invitations/client"
	policies "github.com/hashicorp/terraform-provider-azuread/internal/services/policies/client"
	serviceprincipals "github.com/hashicorp/terraform-provider-azuread/internal/services/serviceprincipals/client"
	users "github.com/hashicorp/terraform-provider-azuread/internal/services/users/client"
)

// Client contains the handles to all the specific Azure AD resource classes' respective clients
type Client struct {
	Environment environments.Environment
	TenantID    string
	ClientID    string
	Claims      auth.Claims

	TerraformVersion string

	StopContext context.Context

	AdministrativeUnits *administrativeunits.Client
	Applications        *applications.Client
	AppRoleAssignments  *approleassignments.Client
	ConditionalAccess   *conditionalaccess.Client
	DirectoryRoles      *directoryroles.Client
	Domains             *domains.Client
	Groups              *groups.Client
	Invitations         *invitations.Client
	Policies            *policies.Client
	ServicePrincipals   *serviceprincipals.Client
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
	client.Invitations = invitations.NewClient(o)
	client.Policies = policies.NewClient(o)
	client.ServicePrincipals = serviceprincipals.NewClient(o)
	client.Users = users.NewClient(o)

	// Acquire an access token upfront, so we can decode the JWT and populate the claims
	token, err := o.Authorizer.Token()
	if err != nil {
		return fmt.Errorf("unable to obtain access token: %v", err)
	}
	client.Claims, err = auth.ParseClaims(token)
	if err != nil {
		return fmt.Errorf("unable to parse claims in access token: %v", err)
	}

	// Log the claims for debugging
	claimsJson, err := json.Marshal(client.Claims)
	if err != nil {
		log.Printf("[DEBUG] AzureAD Provider could not marshal access token claims for log outout")
	} else if claimsJson == nil {
		log.Printf("[DEBUG] AzureAD Provider marshaled access token claims was nil")
	} else {
		log.Printf("[DEBUG] AzureAD Provider access token claims: %s", claimsJson)
	}

	// Missing object ID of token holder will break many things
	if client.Claims.ObjectId == "" {
		return fmt.Errorf("parsing claims in access token: oid claim is empty")
	}

	return nil
}
