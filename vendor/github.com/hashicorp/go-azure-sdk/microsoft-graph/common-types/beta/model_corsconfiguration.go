package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CorsConfiguration struct {
	// The request headers that the origin domain may specify on the CORS request. The wildcard character * indicates that
	// any header beginning with the specified prefix is allowed.
	AllowedHeaders *[]string `json:"allowedHeaders,omitempty"`

	// The HTTP request methods that the origin domain may use for a CORS request.
	AllowedMethods *[]string `json:"allowedMethods,omitempty"`

	// The origin domains that are permitted to make a request against the service via CORS. The origin domain is the domain
	// from which the request originates. The origin must be an exact case-sensitive match with the origin that the user age
	// sends to the service.
	AllowedOrigins *[]string `json:"allowedOrigins,omitempty"`

	// The maximum amount of time that a browser should cache the response to the preflight OPTIONS request.
	MaxAgeInSeconds nullable.Type[int64] `json:"maxAgeInSeconds,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Resource within the application segment for which CORS permissions are granted. / grants permission for whole app
	// segment.
	Resource nullable.Type[string] `json:"resource,omitempty"`
}
