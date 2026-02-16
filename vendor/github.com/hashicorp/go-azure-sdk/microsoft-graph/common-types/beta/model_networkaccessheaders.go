package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessHeaders struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents the origin or source from which the request is being made.
	Origin nullable.Type[string] `json:"origin,omitempty"`

	// Represents the referring URL or the URL of the web page that the current request originates from.
	Referrer nullable.Type[string] `json:"referrer,omitempty"`

	// Represents the information about the client original IP address when the request passes through one or more proxy
	// servers or load balancers.
	XForwardedFor nullable.Type[string] `json:"xForwardedFor,omitempty"`
}
