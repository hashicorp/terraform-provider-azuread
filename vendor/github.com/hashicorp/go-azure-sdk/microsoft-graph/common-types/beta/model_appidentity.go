package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppIdentity struct {
	// Refers to the unique identifier representing application ID in the Microsoft Entra ID.
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// Refers to the application name displayed in the Microsoft Entra admin center.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Refers to the unique identifier indicating service principal ID in Microsoft Entra ID for the corresponding App.
	ServicePrincipalId nullable.Type[string] `json:"servicePrincipalId,omitempty"`

	// Refers to the Service Principal Name is the Application name in the tenant.
	ServicePrincipalName nullable.Type[string] `json:"servicePrincipalName,omitempty"`
}
