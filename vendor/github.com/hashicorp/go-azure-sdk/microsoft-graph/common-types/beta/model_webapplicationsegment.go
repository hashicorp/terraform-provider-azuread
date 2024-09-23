package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ApplicationSegment = WebApplicationSegment{}

type WebApplicationSegment struct {
	// If you're configuring a traffic manager in front of multiple App Proxy application segments, this property contains
	// the user-friendly URL that will point to the traffic manager.
	AlternateUrl nullable.Type[string] `json:"alternateUrl,omitempty"`

	// A collection of CORS Rule definitions for a particular application segment.
	CorsConfigurations *[]CorsConfigurationv2 `json:"corsConfigurations,omitempty"`

	// The published external URL for the application segment; for example, https://intranet.contoso.com/.
	ExternalUrl nullable.Type[string] `json:"externalUrl,omitempty"`

	// The internal URL of the application segment; for example, https://intranet/.
	InternalUrl nullable.Type[string] `json:"internalUrl,omitempty"`

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

func (s WebApplicationSegment) ApplicationSegment() BaseApplicationSegmentImpl {
	return BaseApplicationSegmentImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s WebApplicationSegment) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WebApplicationSegment{}

func (s WebApplicationSegment) MarshalJSON() ([]byte, error) {
	type wrapper WebApplicationSegment
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WebApplicationSegment: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WebApplicationSegment: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.webApplicationSegment"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WebApplicationSegment: %+v", err)
	}

	return encoded, nil
}
