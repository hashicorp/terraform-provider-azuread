package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OptionalClaim struct {
	// Additional properties of the claim. If a property exists in this collection, it modifies the behavior of the optional
	// claim specified in the name property.
	AdditionalProperties *[]string `json:"additionalProperties,omitempty"`

	// If the value is true, the claim specified by the client is necessary to ensure a smooth authorization experience for
	// the specific task requested by the end user. The default value is false.
	Essential *bool `json:"essential,omitempty"`

	// The name of the optional claim.
	Name *string `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The source (directory object) of the claim. There are predefined claims and user-defined claims from extension
	// properties. If the source value is null, the claim is a predefined optional claim. If the source value is user, the
	// value in the name property is the extension property from the user object.
	Source nullable.Type[string] `json:"source,omitempty"`
}
