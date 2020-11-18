package clients

import (
	"context"

	"github.com/Azure/go-autorest/autorest/azure"

	aad "github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/client"
)

// AadClient contains the handles to all the specific Azure AD resource classes' respective clients
type AadClient struct {
	// todo move this to an "Account" struct as in azurerm?
	ClientID         string
	ObjectID         string
	TenantID         string
	TerraformVersion string
	Environment      azure.Environment

	AuthenticatedAsAServicePrincipal bool

	StopContext context.Context

	// Azure AD clients
	AadGraph *aad.Client
}
