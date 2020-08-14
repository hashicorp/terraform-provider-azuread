package clients

import (
	"context"

	"github.com/Azure/go-autorest/autorest/azure"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/services"

	aad "github.com/terraform-providers/terraform-provider-azuread/azuread/internal/services/aadgraph/client"
)

// AadClient contains the handles to all the specific Azure AD resource classes' respective clients
type AadClient struct {
	// todo move this to an "Account" struct as in azurerm?
	ClientID         string
	ObjectID         string
	SubscriptionID   string
	TenantID         string
	TerraformVersion string
	Environment      azure.Environment

	AuthenticatedAsAServicePrincipal bool

	StopContext context.Context

	// Azure AD clients
	AadGraph *aad.Client
}
