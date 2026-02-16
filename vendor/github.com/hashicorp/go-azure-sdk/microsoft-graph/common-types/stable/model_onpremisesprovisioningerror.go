package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesProvisioningError struct {
	// Category of the provisioning error. Note: Currently, there is only one possible value. Possible value:
	// PropertyConflict - indicates a property value is not unique. Other objects contain the same value for the property.
	Category nullable.Type[string] `json:"category,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The date and time at which the error occurred.
	OccurredDateTime nullable.Type[string] `json:"occurredDateTime,omitempty"`

	// Name of the directory property causing the error. Current possible values: UserPrincipalName or ProxyAddress
	PropertyCausingError nullable.Type[string] `json:"propertyCausingError,omitempty"`

	// Value of the property causing the error.
	Value nullable.Type[string] `json:"value,omitempty"`
}
