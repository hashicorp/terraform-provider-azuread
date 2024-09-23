package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesApplicationSegment struct {
	// If you're configuring a traffic manager in front of multiple App Proxy application segments, contains the
	// user-friendly URL that will point to the traffic manager.
	AlternateUrl nullable.Type[string] `json:"alternateUrl,omitempty"`

	// CORS Rule definition for a particular application segment.
	CorsConfigurations *[]CorsConfiguration `json:"corsConfigurations,omitempty"`

	// The published external URL for the application segment; for example, https://intranet.contoso.com./
	ExternalUrl nullable.Type[string] `json:"externalUrl,omitempty"`

	// The internal URL of the application segment; for example, https://intranet/.
	InternalUrl nullable.Type[string] `json:"internalUrl,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
