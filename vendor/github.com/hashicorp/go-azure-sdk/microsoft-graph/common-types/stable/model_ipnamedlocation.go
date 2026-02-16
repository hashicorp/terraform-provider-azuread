package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NamedLocation = IPNamedLocation{}

type IPNamedLocation struct {
	// List of IP address ranges in IPv4 CIDR format (for example, 1.2.3.4/32) or any allowable IPv6 format from IETF
	// RFC5969. Required.
	IPRanges []IPRange `json:"ipRanges"`

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

func (s IPNamedLocation) NamedLocation() BaseNamedLocationImpl {
	return BaseNamedLocationImpl{
		CreatedDateTime:  s.CreatedDateTime,
		DisplayName:      s.DisplayName,
		ModifiedDateTime: s.ModifiedDateTime,
		Id:               s.Id,
		ODataId:          s.ODataId,
		ODataType:        s.ODataType,
	}
}

func (s IPNamedLocation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IPNamedLocation{}

func (s IPNamedLocation) MarshalJSON() ([]byte, error) {
	type wrapper IPNamedLocation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IPNamedLocation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IPNamedLocation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ipNamedLocation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IPNamedLocation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IPNamedLocation{}

func (s *IPNamedLocation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IsTrusted        *bool                 `json:"isTrusted,omitempty"`
		CreatedDateTime  nullable.Type[string] `json:"createdDateTime,omitempty"`
		DisplayName      *string               `json:"displayName,omitempty"`
		ModifiedDateTime nullable.Type[string] `json:"modifiedDateTime,omitempty"`
		Id               *string               `json:"id,omitempty"`
		ODataId          *string               `json:"@odata.id,omitempty"`
		ODataType        *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IsTrusted = decoded.IsTrusted
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.Id = decoded.Id
	s.ModifiedDateTime = decoded.ModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IPNamedLocation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["ipRanges"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling IPRanges into list []json.RawMessage: %+v", err)
		}

		output := make([]IPRange, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIPRangeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'IPRanges' for 'IPNamedLocation': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.IPRanges = output
	}

	return nil
}
