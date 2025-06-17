package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationBehaviors struct {
	// If false, allows the app to have extended access to Azure AD Graph until June 30, 2025 when Azure AD Graph is fully
	// retired. For more information on Azure AD retirement updates, see June 2024 update on Azure AD Graph API retirement.
	BlockAzureADGraphAccess nullable.Type[bool] `json:"blockAzureADGraphAccess,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// If true, removes the email claim from tokens sent to an application when the email address's domain can't be
	// verified.
	RemoveUnverifiedEmailClaim nullable.Type[bool] `json:"removeUnverifiedEmailClaim,omitempty"`

	// If true, requires multitenant applications to have a service principal in the resource tenant as part of
	// authorization checks before they're granted access tokens. This property is only modifiable for multitenant resource
	// applications that rely on access from clients without a service principal and had this behavior as set to false by
	// Microsoft. Tenant administrators should respond to security advisories sent through Azure Health Service events and
	// the Microsoft 365 message center.
	RequireClientServicePrincipal nullable.Type[bool] `json:"requireClientServicePrincipal,omitempty"`
}
