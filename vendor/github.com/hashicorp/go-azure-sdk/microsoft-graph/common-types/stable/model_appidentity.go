package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppIdentity struct {
	// Refers to the unique ID representing application in Microsoft Entra ID.
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// Refers to the application name displayed in the Microsoft Entra admin center.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Refers to the unique ID for the service principal in Microsoft Entra ID.
	ServicePrincipalId nullable.Type[string] `json:"servicePrincipalId,omitempty"`

	// Refers to the Service Principal Name is the Application name in the tenant.
	ServicePrincipalName nullable.Type[string] `json:"servicePrincipalName,omitempty"`
}
