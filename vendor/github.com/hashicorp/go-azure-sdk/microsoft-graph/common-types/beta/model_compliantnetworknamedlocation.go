package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NamedLocation = CompliantNetworkNamedLocation{}

type CompliantNetworkNamedLocation struct {
	CompliantNetworkType *CompliantNetworkType `json:"compliantNetworkType,omitempty"`

	// true if this location is explicitly trusted. Optional. Default value is false.
	IsTrusted *bool `json:"isTrusted,omitempty"`

	// Fields inherited from NamedLocation

	// The Timestamp type represents creation date and time of the location using ISO 8601 format and is always in UTC time.
	// For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Human-readable name of the location.
	DisplayName *string `json:"displayName,omitempty"`

	// The Timestamp type represents last modified date and time of the location using ISO 8601 format and is always in UTC
	// time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	ModifiedDateTime nullable.Type[string] `json:"modifiedDateTime,omitempty"`

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

func (s CompliantNetworkNamedLocation) NamedLocation() BaseNamedLocationImpl {
	return BaseNamedLocationImpl{
		CreatedDateTime:  s.CreatedDateTime,
		DisplayName:      s.DisplayName,
		ModifiedDateTime: s.ModifiedDateTime,
		Id:               s.Id,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

func (s CompliantNetworkNamedLocation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CompliantNetworkNamedLocation{}

func (s CompliantNetworkNamedLocation) MarshalJSON() ([]byte, error) {
	type wrapper CompliantNetworkNamedLocation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CompliantNetworkNamedLocation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CompliantNetworkNamedLocation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.compliantNetworkNamedLocation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CompliantNetworkNamedLocation: %+v", err)
	}

	return encoded, nil
}
