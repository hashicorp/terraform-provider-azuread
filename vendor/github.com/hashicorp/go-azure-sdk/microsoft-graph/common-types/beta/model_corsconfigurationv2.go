package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CorsConfigurationv2{}

type CorsConfigurationv2 struct {
	// The request headers that the origin domain may specify on the CORS request. The wildcard character * indicates that
	// any header beginning with the specified prefix is allowed.
	AllowedHeaders *[]string `json:"allowedHeaders,omitempty"`

	// The HTTP request methods that the origin domain may use for a CORS request.
	AllowedMethods *[]string `json:"allowedMethods,omitempty"`

	// The origin domains that are permitted to make a request against the service via CORS. The origin domain is the domain
	// from which the request originates. The origin must be an exact case-sensitive match with the origin that the user
	// agent sends to the service.
	AllowedOrigins *[]string `json:"allowedOrigins,omitempty"`

	// The maximum amount of time that a browser should cache the response to the preflight OPTIONS request.
	MaxAgeInSeconds nullable.Type[int64] `json:"maxAgeInSeconds,omitempty"`

	// Resource within the application segment for which CORS permissions are granted. / grants permission for the whole app
	// segment.
	Resource *string `json:"resource,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CorsConfigurationv2) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CorsConfigurationv2{}

func (s CorsConfigurationv2) MarshalJSON() ([]byte, error) {
	type wrapper CorsConfigurationv2
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CorsConfigurationv2: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CorsConfigurationv2: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.corsConfiguration_v2"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CorsConfigurationv2: %+v", err)
	}

	return encoded, nil
}
