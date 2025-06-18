package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsConfiguration struct {
	// A collection of application IDs for registered Microsoft Entra apps allowed to manage the externalConnection and
	// index content in the externalConnection.
	AuthorizedAppIds *[]string `json:"authorizedAppIds,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
